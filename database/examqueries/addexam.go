package examqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddExam inserta en la base de datos el modelo de exam
func AddExam(examModel models.Exam) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Exam")

	result, err := col.InsertOne(ctx, examModel)

	if err != nil {
		return "", false, err
	}
	id, _ := result.InsertedID.(primitive.ObjectID)

	return id.Hex(), true, nil
}
