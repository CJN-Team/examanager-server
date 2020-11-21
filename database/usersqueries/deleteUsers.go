package usersqueries

import (
	"context"
	"errors"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"go.mongodb.org/mongo-driver/bson"
)

//DeleteUser se encarga de borrar el usuario seleccionado
func DeleteUser(ID string, loggedUser string, loggedInstitution string) error {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	userID := ID

	condicion := bson.M{
		"_id": userID,
		"institution":loggedInstitution,
	}

	if userTypeVerificationdeleting(loggedUser, loggedInstitution) {
		error := errors.New("el usuario no posee los permisos suficientes")
		return error
	}

	_, error := coleccion.DeleteOne(contex, condicion)

	return error
}

func userTypeVerificationdeleting(loggedUser string, loggedInstitution string) bool {

	userID, _ := GetUserByIDOneInstitution(loggedUser, loggedInstitution)

	if userID.Profile != "Administrador" {
		return true
	}
	return false
}
