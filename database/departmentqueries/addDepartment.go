package departmentqueries

import (
	"context"
	"errors"
	"fmt"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/database/usersqueries"
	"github.com/CJN-Team/examanager-server/models"
)

//AddDepartment se encarga de añadir a la base de datos un nuevo departamento
func AddDepartment(department models.Department, loggedUser string, institucionID string) (string, bool, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("departments")

	if loggedUser != "" {
		_, admin := userTypeVerificationAdding(loggedUser)
		if admin {
			error := errors.New("el usuario no posee los permisos suficientes")
			return "", false, error
		}
	}

	errorUsers := verifyIfTeachersExist(department.Teachers)

	if errorUsers != "" {
		error := errors.New(errorUsers)
		return "", false, error
	}

	_, error := coleccion.InsertOne(contex, department)

	if error != nil {
		return "", false, error
	}

	return "", true, nil
}

func userTypeVerificationAdding(loggedUser string) (string, bool) {

	userID, _ := usersqueries.GetUserByID(loggedUser)

	if userID.Profile != "Administrador" {
		return "", true
	}
	return userID.Institution, false
}

func verifyIfTeachersExist(users []string) string {

	errors := false
	wrongUsers := "Usuarios invalidos o no registrados: \n"
	for _, user := range users {

		userModel, error := usersqueries.GetUserByID(fmt.Sprintf("%v", user))

		if error != nil || userModel.Profile != "Profesor" {
			wrongUsers += fmt.Sprintf("%v", user) + "\n"
			errors = true
		}
	}

	if errors {
		return wrongUsers
	}
	return ""

}
