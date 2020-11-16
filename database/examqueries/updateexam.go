package examqueries

import (
	"context"
	"fmt"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UpdateExam se encarga de actualizar la informacion de la pregunta
func UpdateExam(exam models.Exam, ID string) (bool, error) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("Exam")

	examRegisterd := make(map[string]interface{})

	if len(exam.SubjectID) > 0 {
		examRegisterd["subjectID"] = exam.SubjectID
	}

	if len(exam.GroupID) > 0 {
		examRegisterd["groupID"] = exam.GroupID
	}

	if len(exam.Name) > 0 {
		examRegisterd["name"] = exam.Name
	}

	if exam.State {
		examRegisterd["state"] = exam.State
	} else {
		examRegisterd["state"] = exam.State
	}

	if len(exam.Difficulty) > 0 {
		examRegisterd["difficulty"] = exam.Difficulty
	}

	examRegisterd["date"] = exam.Date

	if len(exam.TopicQuestion) > 0 {
		examRegisterd["topicQuestion"] = exam.TopicQuestion
	}

	if len(exam.GenerateExam) > 0 {
		examRegisterd["generateExam"] = exam.GenerateExam
	}

	updateString := bson.M{
		"$set": examRegisterd,
	}

	id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": id}}

	a, error := coleccion.UpdateOne(contex, filter, updateString)

	fmt.Println("eso ", a)

	if error != nil {
		return false, error
	}

	return true, nil
}

//UpdateExamGrade actualiza la nota de un examen dado el ID del examen
func UpdateExamGrade(examID string, grade float32)(bool, error){

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("GenerateExam")

	updateString := bson.M{
		"$set": bson.M{
			"grade" : grade,
		},
	}

	id, _ := primitive.ObjectIDFromHex(examID )
	filter := bson.M{"_id": bson.M{"$eq": id}}

	_, error := coleccion.UpdateOne(contex, filter, updateString)
	
	if error != nil {
		return false, error
	}
	return true, nil
}