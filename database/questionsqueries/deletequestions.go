package questionsqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"go.mongodb.org/mongo-driver/bson"
)

//DeleteQuestion elimina una pregunta de la base de datos
func DeleteQuestion(ID string) error {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("Questions")

	condicion := bson.M{
		"_id": ID,
	}

	_, error := coleccion.DeleteOne(contex, condicion)

	return error
}
