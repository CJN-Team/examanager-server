package generatexamqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
