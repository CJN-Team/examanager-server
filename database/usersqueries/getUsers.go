package usersqueries

import (
	"context"
	"log"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllUsers se encarga de traer de base de datos todos los usuarios disponibles de una categoria
func GetAllUsers(category string, institution string, page int64) ([]*models.User, bool) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	var result []*models.User

	condicion := bson.M{
		"profile":     category,
		"institution": institution,
	}

	searchOptions := options.Find()
	searchOptions.SetLimit(20)
	searchOptions.SetSort(bson.D{{Key: "name", Value: -1}})
	searchOptions.SetSkip((page - 1) * 20)

	pointer, error := coleccion.Find(contex, condicion, searchOptions)

	if error != nil {
		log.Fatal(error.Error())
		return result, false
	}

	for pointer.Next(context.TODO()) {
		var register models.User
		error := pointer.Decode(&register)

		if error != nil {
			return result, false
		}
		result = append(result, &register)
	}

	return result, true
}

//GetUserByEmail verifica si el usuario ya se encuentra en la base de datos por medio de el correo
func GetUserByEmail(email string) (models.User, bool, string) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	condicion := bson.M{"email": email}

	var result models.User

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	ID := result.ID

	if error != nil {
		return result, false, ID
	}

	return result, true, ID
}

//GetUserByIDAllInstitutions se encarga de buscar en la base de datos el usuario que posee la ID asignada
func GetUserByIDAllInstitutions(ID string) (models.User, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	var result models.User

	ObjectID := ID

	condicion := bson.M{"_id": ObjectID}

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	result.Password = ""

	if error != nil {
		return result, error
	}

	return result, nil
}

//GetUserByIDOneInstitution se encarga de buscar en la base de datos el usuario que posee la ID asignada en una sola institucion
func GetUserByIDOneInstitution(ID string, institution string) (models.User, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	var result models.User

	ObjectID := ID

	condicion := bson.M{"_id": ObjectID, "institution":institution}

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	result.Password = ""

	if error != nil {
		return result, error
	}

	return result, nil
}