package database

import (
	"context"
	"time"

	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddRegister se encarga de añadir a la base de datos un nuevo usuario
func AddRegister(u models.User) (string, bool, error) {

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
