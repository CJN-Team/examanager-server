package institutionsqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddSubject le permite a un administrador de una institucion crear una asignatura
func AddSubject(institutionInfo models.Institution) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Institutions")

	updateString := bson.M{
		"$set": institutionInfo,
	}

	id, _ := primitive.ObjectIDFromHex(institutionInfo.ID.Hex())
	filter := bson.M{"_id": bson.M{"$eq": id}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil

}

//DeleteSubject elimina una asignatura de una institucion
func DeleteSubject(institutionInfo models.Institution, SubjectName string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Institutions")
	delete(institutionInfo.Subjetcs, SubjectName)

	updateString := bson.M{
		"$set": institutionInfo,
	}

	id, _ := primitive.ObjectIDFromHex(institutionInfo.ID.Hex())
	filter := bson.M{"_id": bson.M{"$eq": id}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil

}

/*func GetSubject(institutionInfo models.Institution, SubjectName string)(bool,error){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Institutions")

	id, _ := primitive.ObjectIDFromHex(institutionInfo.ID.Hex())
	filter := bson.M{"_id": bson.M{"$eq": id}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil

}*/
