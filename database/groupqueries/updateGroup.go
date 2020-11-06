package groupqueries

import (
	"context"
	"errors"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/database/usersqueries"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UpdateGroup se encarga de actualizar el usuario registrado
func UpdateGroup(group models.Group, ID string, loggedUser string) (bool, error) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("groups")

	groupRegisterd := make(map[string]interface{})

	if len(group.StudentsList) > 0 {
		groupRegisterd["studentsList"] = group.StudentsList
	}

	if len(group.Teacher) > 0 {
		groupRegisterd["teacher"] = group.Teacher
	}

	if len(group.Subject) > 0 {
		groupRegisterd["subject"] = group.Subject
	}

	updateString := bson.M{
		"$set": groupRegisterd,
	}

	if userTypeVerificationdeleting(loggedUser) {
		error := errors.New("el usuario no posee los permisos suficientes")
		return false, error
	}

	filter := bson.M{"_id": bson.M{"$eq": ID}}

	_, error := coleccion.UpdateOne(contex, filter, updateString)

	if error != nil {
		return false, error
	}

	return true, nil
}

func userTypeVerificationUpdating(loggedUser string) bool {

	userID, _ := usersqueries.GetUserByID(loggedUser)

	if userID.Profile != "Administrador" {
		return true
	}
	return false
}
