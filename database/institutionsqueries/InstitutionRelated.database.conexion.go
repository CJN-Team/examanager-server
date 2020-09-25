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
func AddSubject(institutionInfo models.Institution)(bool,error){

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
func DeleteSubject(institutionInfo models.Institution, SubjectName string)(bool,error){

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



//AddUsersXInstitution crea un documento que relacionara la institucion que se esta creando con los diferentes usuarios de esta.
func AddUsersXInstitution(name string)(string,bool,error){

	var UsersXInstitutionModel models.UsersXInstitution
	UsersXInstitutionModel.InstitutionName=name

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("UsersXInstitution")

	result, err := col.InsertOne(ctx, UsersXInstitutionModel)
	if err != nil {
		return "", false, err
	}

	UsersXInstitutionID, _ := result.InsertedID.(primitive.ObjectID)
	return UsersXInstitutionID.Hex(), true, nil

}
//AddQuestionsXInstitution crea un documento que relacionara la institucion que se esta creando con las preguntas de esta.
func AddQuestionsXInstitution(name string)(string,bool,error){

	var QuestionsXInstitutionModel models.QuestionsXInstitution
	QuestionsXInstitutionModel.InstitutionName=name

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("QuestionsXInstitution")

	result, err := col.InsertOne(ctx, QuestionsXInstitutionModel)
	if err != nil {
		return "", false, err
	}

	QuestionsXInstitutionID, _ := result.InsertedID.(primitive.ObjectID)
	return QuestionsXInstitutionID.Hex(), true, nil

}
//AddQuestionToInstitution añade una pregunta relacionada a una institucion a preguntas X institucion
func AddQuestionToInstitution(questionXInstitutionInfo models.QuestionsXInstitution, name string)(bool,error){


	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("QuestionsXInstitution")

	updateString := bson.M{ "$push": bson.M{"questionsList":name}}

	id,_ := primitive.ObjectIDFromHex(questionXInstitutionInfo.ID.Hex())

	filter := bson.M{"_id": bson.M{"$eq": id}}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil

}
