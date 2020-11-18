package examqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UpdateExam se encarga de actualizar la informacion de la pregunta
func UpdateExam(exam models.Exam, ID string) (bool, error) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("Exam")

	examRegisterd := make(map[string]interface{})

	if len(exam.SubjectID) > 0 {
		examRegisterd["subjectID"] = exam.SubjectID
	}

	if len(exam.GroupID) > 0 {
		examRegisterd["groupId"] = exam.GroupID
	}

	if len(exam.Name) > 0 {
		examRegisterd["name"] = exam.Name
	}

	examRegisterd["state"] = exam.State

	examRegisterd["view"] = exam.View

	if len(exam.Difficulty) > 0 {
		examRegisterd["difficulty"] = exam.Difficulty
	}

	examRegisterd["date"] = exam.Date

	if len(exam.TopicQuestion) > 0 {
		examRegisterd["topicQuestion"] = exam.TopicQuestion
	}

	if len(exam.GenerateExam) > 0 {
		examRegisterd["generateExam"] = exam.GenerateExam
	}

	updateString := bson.M{
		"$set": examRegisterd,
	}

	id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": id}}

	_, error := coleccion.UpdateOne(contex, filter, updateString)

	if error != nil {
		return false, error
	}

	return true, nil
}

//UpdateGenerateExam actualiza algunos campos del examen generado
func UpdateGenerateExam(exam models.GenerateExam, ID string) (bool, error) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("GenerateExam")
	examRegisterd := make(map[string]interface{})

	examRegisterd["grade"] = exam.Grade

	examRegisterd["finish"] = exam.Finish

	if len(exam.Commentary) > 0 {
		examRegisterd["commentary"] = exam.Commentary
	}

	updateString := bson.M{
		"$set": examRegisterd,
	}

	id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": id}}

	_, error := coleccion.UpdateOne(contex, filter, updateString)

	if error != nil {
		return false, error
	}
	return true, nil
}
