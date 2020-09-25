package questionsqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeleteQuestion elimina una pregunta de la base de datos
func DeleteQuestion(ID string) error {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("Questions")

	userID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": userID,
	}

	_, error := coleccion.DeleteOne(contex, condicion)

	return error
}
