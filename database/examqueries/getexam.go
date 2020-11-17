package examqueries

import (
	"context"
	"log"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetExamByID verifica si el examen ya se encuentra en la base de datos por medio del id
func GetExamByID(id string, institution string) (models.Exam, bool, string) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")
	coleccion := database.Collection("Exam")
	var result models.Exam
	idaux, _ := primitive.ObjectIDFromHex(id)
	error := coleccion.FindOne(contex, bson.M{"_id": idaux, "institution": institution}).Decode(&result)

	ID := result.ID

	if error != nil {
		return result, false, ID.Hex()
	}

	return result, true, ID.Hex()
}

//GetExamByName verifica si el examen ya existe por nombre
func GetExamByName(name string, group string, institution string) (models.Exam, bool, string) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")
	coleccion := database.Collection("Exam")

	condicion := bson.M{"name": name, "groupId":group, "institution": institution}

	var result models.Exam

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	ID := result.ID

	if error != nil {
		return result, false, ID.Hex()
	}

	return result, true, ID.Hex()
}

//GetAllExamByGroup trae todos los examenes por grupo
func GetAllExamByGroup(groupID string, institution string, page int64) ([]*models.Exam, bool) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")
	coleccion := database.Collection("Exam")

	condicion := bson.M{"groupId": groupID, "institution": institution}

	var result []*models.Exam

	searchOptions := options.Find()
	searchOptions.SetLimit(20)
	searchOptions.SetSkip((page - 1) * 20)

	pointer, error := coleccion.Find(contex, condicion, searchOptions)

	if error != nil {
		log.Fatal(error.Error())
		return result, false
	}

	for pointer.Next(context.TODO()) {
		var register models.Exam
		error := pointer.Decode(&register)

		if error != nil {
			return result, false
		}
		result = append(result, &register)
	}

	return result, true
}
