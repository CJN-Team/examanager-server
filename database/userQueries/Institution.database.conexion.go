package userqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddInstitution inserta en la base de datos el modelo de la institución creada
func AddInstitution(institutionModel models.Institution) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Institutions")

	result, err := col.InsertOne(ctx, institutionModel)
	if err != nil {
		return "", false, err
	}

	InstitutionID, _ := result.InsertedID.(primitive.ObjectID)
	return InstitutionID.String(), true, nil
}

//GetInstitution busca en la base de datos la existencia de una institucion por el nombre
func GetInstitution(name string) (models.Institution, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Institutions")
	filter := bson.M{"name": name}

	var institution models.Institution
	err := col.FindOne(ctx, filter).Decode(&institution)
	institutionID := institution.ID.Hex()
	if err != nil {
		return institution, false, institutionID
	}
	return institution, true, institutionID
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
