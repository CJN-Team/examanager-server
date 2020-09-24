package institutionqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
)

func AddSubject(SubjectInfo models.Subject, institutionInfo models.Institution)(bool,error){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Institutions")

	fmt.Printf("%v \n",SubjectInfo)
	fmt.Printf("%v \n",institutionInfo)
	//institutionInfo.Subjetcs = append(institutionInfo.Subjetcs,SubjectInfo)
	fmt.Printf("%v \n",institutionInfo.Subjetcs)
	updateString := bson.M{
		"$set": institutionInfo,
	}
	fmt.Printf("%v \n",updateString)
	filter := bson.M{"_id": bson.M{"$eq": institutionInfo.ID.Hex()}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil

}


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
	return UsersXInstitutionID.String(), true, nil

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
	return QuestionsXInstitutionID.String(), true, nil

}

