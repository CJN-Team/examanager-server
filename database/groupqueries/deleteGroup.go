package groupqueries

import (
	"context"
	"errors"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/database/usersqueries"
	"go.mongodb.org/mongo-driver/bson"
)

//DeleteGroup se encarga de borrar el grupo seleccionado
func DeleteGroup(ID string, loggedUser string) error {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("groups")

	groupID := ID

	condicion := bson.M{
		"_id": groupID,
	}

	if userTypeVerificationdeleting(loggedUser) {
		error := errors.New("el usuario no posee los permisos suficientes")
		return error
	}

	exist, error := coleccion.DeleteOne(contex, condicion)
	if exist.DeletedCount == 0 {
		error := errors.New("el archivo a eliminar no existe")
		return error
	}
	return error
}

func userTypeVerificationdeleting(loggedUser string) bool {

	userID, _ := usersqueries.GetUserByID(loggedUser)

	if userID.Profile != "Administrador" {
		return true
	}
	return false
}
