package institutionsqueries

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

//GetInstitutionByName busca en la base de datos la existencia de una institucion por el nombre
func GetInstitutionByName(name string) (models.Institution, bool, string) {
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

//GetInstitutionByID busca en la base de datos la existencia de una institucion por el nombre
func GetInstitutionByID(InstitutionID string) (models.Institution, bool,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Institutions")	

	var institutionInfo models.Institution
	id,_ := primitive.ObjectIDFromHex(InstitutionID)
	err := col.FindOne(ctx,bson.M{"_id": id}).Decode(&institutionInfo)

	if err != nil {
		return institutionInfo, false,err
	}
	return institutionInfo, true,nil
}
