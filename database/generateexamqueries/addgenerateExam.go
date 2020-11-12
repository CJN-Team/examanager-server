package generatexamqueries

import (
	"context"
	"math/rand"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	grupDB "github.com/CJN-Team/examanager-server/database/groupqueries"
	institutionDB "github.com/CJN-Team/examanager-server/database/institutionsqueries"
	questionsDB "github.com/CJN-Team/examanager-server/database/questionsqueries"
	userDB "github.com/CJN-Team/examanager-server/database/usersqueries"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddGenerateExam inserta en la base de datos el modelo de examenes generados para los estudiantes
func AddGenerateExam(generateExamModel models.GenerateExam) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("GenerateExam")

	result, err := col.InsertOne(ctx, generateExamModel)

	if err != nil {
		return "", false, err
	}
	id, _ := result.InsertedID.(primitive.ObjectID)

	return id.Hex(), true, nil
}

//GenerateExam genera los examenes
func GenerateExam(examModel models.Exam, loggedUser string, institution string) ([]string, bool, error) {
	var generateExam models.GenerateExam
	var ids []string
	group, err := grupDB.GetGroupByID(examModel.GroupID, institution)

	if err != nil {
		return ids, false, err
	}

	for key, value := range group.StudentsList {
		generateExam.Teacher = group.Teacher
		student, _ := userDB.GetUserByID(key)
		generateExam.Student = student.Name + " " + student.LastName
		generateExam.Date = examModel.Date
		institution, _, _ := institutionDB.GetInstitutionByID(student.Institution)
		generateExam.Institution = institution.Name
		generateExam.Photo = student.Photo
		generateExam.Name = examModel.Name
		generateExam.Question, _, _ = GetQuestions(examModel, student.Institution)

		id, _, _ := AddGenerateExam(generateExam)

		group.StudentsList[key] = append(value.(primitive.A), id)

		ids = append(ids, id)
	}
	grupDB.UpdateGroup(group, group.ID, loggedUser)
	return ids, true, nil
}

//GetQuestions trae las preguntas necesarias para el examen
func GetQuestions(examModel models.Exam, institution string) ([]string, bool, error) {
	var questions []string
	var random int
	questionsFacil, _ := questionsDB.GetAllQuestions(examModel.TopicQuestion, "1", 3, -1, institution)
	questionsNormal, _ := questionsDB.GetAllQuestions(examModel.TopicQuestion, "2", 3, -1, institution)
	questionsDificil, _ := questionsDB.GetAllQuestions(examModel.TopicQuestion, "3", 3, -1, institution)
	facil := examModel.Difficulty[0]
	normal := examModel.Difficulty[1]
	dificil := examModel.Difficulty[2]

	i := 0
	for i < facil {
		random = rand.Intn(len(questionsFacil) - 1)
		if !contains(questions, questionsFacil[random].ID) {
			questions = append(questions, questionsFacil[random].ID)
			i++
		}
	}
	i = 0
	for i < normal {
		random = rand.Intn(len(questionsNormal) - 1)
		if !contains(questions, questionsNormal[random].ID) {
			questions = append(questions, questionsNormal[random].ID)
			i++
		}
	}
	i = 0
	for i < dificil {
		random = rand.Intn(len(questionsDificil) - 1)
		if !contains(questions, questionsDificil[random].ID) {
			questions = append(questions, questionsDificil[random].ID)
			i++
		}
	}
	return questions, true, nil
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
