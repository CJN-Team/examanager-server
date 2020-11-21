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
		_, admin := userTypeVerificationAdding(loggedUser, institucionID)
		if admin {
			error := errors.New("el usuario no posee los permisos suficientes")
			return "", false, error
		}
	}

	errorUsers := verifyIfTeachersExist(department.Teachers, institucionID)

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

func userTypeVerificationAdding(loggedUser string, loggedUserInstitution string) (string, bool) {

	userID, _ := usersqueries.GetUserByIDOneInstitution(loggedUser, loggedUserInstitution)

	if userID.Profile != "Administrador" {
		return "", true
	}
	return userID.Institution, false
}

func verifyIfTeachersExist(users []string, loggedUserInstitution string) string {

	errors := false
	wrongUsers := "Usuarios invalidos o no registrados: \n"
	for _, user := range users {

		userModel, error := usersqueries.GetUserByIDOneInstitution(fmt.Sprintf("%v", user), loggedUserInstitution)

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
