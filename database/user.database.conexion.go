package database

import (
	"context"
	"time"

	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UserExist verifica si el usuario ya se encuentra en la base de datos por medio de el correo
func UserExist(email string) (models.User, bool, string) {

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

	result, error := coleccion.InsertOne(contex, u)

	if error != nil {
		return "", false, error
	}

	ObjectID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjectID.String(), true, nil
}
