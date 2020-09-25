package questionsqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetQuestionByID busca en la base de datos la existencia de una pregunta por el ID
func GetQuestionByID(QuestionID string) (models.Question, bool,error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Questions")	

	var questionInfo models.Question
	id,_ := primitive.ObjectIDFromHex(QuestionID)
	err := col.FindOne(ctx,bson.M{"_id": id}).Decode(&questionInfo)

	if err != nil {
		return questionInfo, false,err
	}
	return questionInfo, true,nil
}