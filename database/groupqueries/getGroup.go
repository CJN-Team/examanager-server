package groupqueries

import (
	"context"
	"log"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllGroups se encarga de traer de base de datos todos los grupos almacenados
func GetAllGroups(page int64, institution string) ([]*models.Group, bool) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("groups")

	var result []*models.Group

	condicion := bson.M{
		"institution": institution,
	}

	searchOptions := options.Find()
	searchOptions.SetLimit(20)
	searchOptions.SetSkip((page - 1) * 20)

	pointer, error := coleccion.Find(contex, condicion, searchOptions)

	if error != nil {
		log.Fatal(error.Error())
		return result, false
	}

	for pointer.Next(context.TODO()) {
		var register models.Group
		error := pointer.Decode(&register)

		if error != nil {
			return result, false
		}
		result = append(result, &register)
	}

	return result, true
}

//GetGroupByID se encarga de buscar en la base de datos el grupo que posee la ID asignada
func GetGroupByID(ID string, institution string) (models.Group, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("groups")

	var result models.Group

	condicion := bson.M{"_id": ID, "institution": institution}

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	if error != nil {
		return result, error
	}

	return result, nil
}
