package questionsqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UpdateQuestion se encarga de actualizar la informacion de la pregunta
func UpdateQuestion(question models.Question, ID string) (bool, error) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("Questions")

	questionRegisterd := make(map[string]interface{})

	if len(question.Topic) > 0 {
		questionRegisterd["topic"] = question.Topic
	}

	if len(question.Subject) > 0 {
		questionRegisterd["subject"] = question.Subject
	}

	if len(question.Pregunta) > 0 {
		questionRegisterd["question"] = question.Pregunta
	}

	if len(question.Category) > 0 {
		questionRegisterd["category"] = question.Category
	}

	if len(question.Options) > 0 {
		questionRegisterd["options"] = question.Options
	}

	if len(question.Answer) > 0 {
		questionRegisterd["answer"] = question.Answer
	}

	if question.Difficulty != 0 {
		questionRegisterd["difficulty"] = question.Difficulty
	}

	updateString := bson.M{
		"$set": questionRegisterd,
	}

	filter := bson.M{"_id": bson.M{"$eq": ID}}

	_, error := coleccion.UpdateOne(contex, filter, updateString)

	if error != nil {
		return false, error
	}

	return true, nil
}
