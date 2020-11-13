package questionsqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"go.mongodb.org/mongo-driver/bson"
)

//DeleteTest elimina un examen de la base de datos dado su ID
func DeleteTest(testID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Exam")

	condicion := bson.M{
		"_id": testID,
	}

	_, err := col.DeleteOne(ctx, condicion)


	if err != nil {
		return false, err
	}

	return true, nil
}
