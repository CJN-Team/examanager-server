package database

import (
	"context"
	"time"

	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetUserByEmail verifica si el usuario ya se encuentra en la base de datos por medio de el correo
func GetUserByEmail(email string) (models.User, bool, string) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	condicion := bson.M{"email": email}

	var result models.User

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	ID := result.ID.Hex()

	if error != nil {
		return result, false, ID
	}

	return result, true, ID
}

//AddUser se encarga de añadir a la base de datos un nuevo usuario
func AddUser(u models.User) (string, bool, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	u.Password, _ = PasswordEncrypt(u.Password)

	result, error := coleccion.InsertOne(contex, u)

	if error != nil {
		return "", false, error
	}

	ObjectID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjectID.String(), true, nil
}

//GetUserByID se encarga de buscar en la base de datos el usuario que posee la ID asignada
func GetUserByID(ID primitive.ObjectID) (models.User, bool, string) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	condicion := bson.M{"ID": ID}

	var result models.User

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	searchError := result.ID.Hex()
	if error != nil {
		return result, false, searchError
	}

	return result, true, searchError
}
