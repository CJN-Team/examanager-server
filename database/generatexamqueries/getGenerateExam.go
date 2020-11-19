package generatexamqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetGenerateExamByID verifica si el examen ya se encuentra en la base de datos por medio del id
func GetGenerateExamByID(id string, institution string) (models.GenerateExam, bool) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")
	coleccion := database.Collection("GenerateExam")
	var result models.GenerateExam
	idaux, _ := primitive.ObjectIDFromHex(id)
	error := coleccion.FindOne(contex, bson.M{"_id": idaux, "institutionid": institution}).Decode(&result)

	if error != nil {
		return result, false
	}

	return result, true
}

//Cosas de James No tocar

//UserGrades se encarga de mostrar las notas de un alumno
func UserGrades(GroupID string, UserID string, institution string) (map[string]float64, error) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("groups")

	gradesAux := make(map[string]float64)
	var groupModel models.Group

	condicion := bson.M{"_id": GroupID, "institution": institution}

	error := coleccion.FindOne(contex, condicion).Decode(&groupModel)

	for _, value := range groupModel.StudentsList[UserID].(primitive.A) {
		currentExam, _ := GetGenerateExamByID(value.(string), institution)

		gradesAux[currentExam.Name] = currentExam.Grade

	}

	if error != nil {
		return gradesAux, error
	}

	return gradesAux, nil
}

//UserGradesAllGroups se encarga de mostrar las notas de un alumno
func UserGradesAllGroups(UserID string, institution string) (map[string]map[string]float64, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("groups")

	condicion := bson.M{"institution": institution}

	searchOptions := options.Find()

	pointer, error := coleccion.Find(contex, condicion, searchOptions)

	grades := make(map[string]map[string]float64)

	for pointer.Next(context.TODO()) {
		var register models.Group
		error := pointer.Decode(&register)

		if error != nil {
			return grades, error
		}

		if _, ok := register.StudentsList[UserID]; ok {
			gradesAux := make(map[string]float64)
			for _, value := range register.StudentsList[UserID].(primitive.A) {

				currentExam, _ := GetGenerateExamByID(value.(string), institution)

				gradesAux[currentExam.Name] = currentExam.Grade

			}
			grades[register.ID+","+register.Name] = gradesAux
		}

	}

	if error != nil {
		return grades, error
	}

	return grades, nil
}
