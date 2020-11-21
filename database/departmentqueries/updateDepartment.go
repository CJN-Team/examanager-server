package departmentqueries

import (
	"context"
	"errors"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/database/usersqueries"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UpdateDepartment se encarga de actualizar el departamento registrado
func UpdateDepartment(department models.Department, ID string, loggedUser string, loggedUserInstitution string) (bool, error) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("departments")

	departmentRegistered := make(map[string]interface{})

	if len(department.Teachers) > 0 {
		departmentRegistered["teachers"] = department.Teachers
	}

	if len(department.Name) > 0 {
		departmentRegistered["name"] = department.Name
	}

	updateString := bson.M{
		"$set": departmentRegistered,
	}

	if userTypeVerificationUpdating(loggedUser, loggedUserInstitution) {
		error := errors.New("el usuario no posee los permisos suficientes")
		return false, error
	}

	departmentID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": departmentID}, "institution":  bson.M{"$eq": loggedUserInstitution}}

	_, error := coleccion.UpdateOne(contex, filter, updateString)

	if error != nil {
		return false, error
	}

	return true, nil
}

func userTypeVerificationUpdating(loggedUser string, loggedUserInstitution string) bool {

	userID, _ := usersqueries.GetUserByIDOneInstitution(loggedUser, loggedUserInstitution)

	if userID.Profile != "Administrador" {
		return true
	}
	return false
}
