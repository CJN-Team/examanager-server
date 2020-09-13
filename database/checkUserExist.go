package database

import (
	"context"
	"time"

	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UserExist verifica si existe un usuario por medio del correo
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
