package generatexamqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)

//UpdateExam actualiza la informacion de un examen dado su ID
func UpdateExam(examid string, updateString bson.M) (error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("GenerateExam")

	id, _ := primitive.ObjectIDFromHex(examid)
	filter := bson.M{"_id": bson.M{"$eq": id}}
	_, err := col.UpdateOne(ctx, filter, updateString)

	return err
}

