package examqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"go.mongodb.org/mongo-driver/bson"
)

//DeleteExam elimina un examen de la base de datos dado su ID
func DeleteExam(examID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Exam")

	condicion := bson.M{
		"_id": examID,
	}

	_, err := col.DeleteOne(ctx, condicion)


	if err != nil {
		return false, err
	}

	return true, nil
}

//DeleteGeneratedExams elimina los examenes generados de un examen padre
func DeleteGeneratedExams(examID string) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("GenerateExam")

	condicion := bson.M{
		"_id": examID,
	}

	_, err := col.DeleteMany(ctx, condicion)


	if err != nil {
		return false, err
	}

	return true, nil
}

