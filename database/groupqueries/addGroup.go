package groupqueries

import (
	"context"
	"errors"
	"fmt"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/database/institutionsqueries"
	"github.com/CJN-Team/examanager-server/database/usersqueries"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//AddGroup se encarga de añadir a la base de datos un nuevo usuario
func AddGroup(group models.Group, loggedUser string, institucionID string) (string, bool, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("groups")

	if loggedUser != "" {
		_, admin := userTypeVerificationAdding(loggedUser, institucionID)
		if admin {
			error := errors.New("el usuario no posee los permisos suficientes")
			return "", false, error
		}
	}

	if verifyIfTeacherExist(group.Teacher, institucionID) != nil {
		error := errors.New("El profesor es invalido o no esta registrado")
		return "", false, error
	}

	if VerifyIfSubjectExist(group.Subject, institucionID) != "" {
		error := errors.New("La asignatura ingresada es invalida o no existe")
		return "", false, error
	}

	errorUsers := verifyIfStudentExist(group.StudentsList, institucionID)

	if errorUsers != "" {
		error := errors.New(errorUsers)
		return "", false, error
	}

	_, error := coleccion.InsertOne(contex, group)

	if error != nil {
		return "", false, error
	}

	return "", true, nil
}

func verifyIfTeacherExist(teacher string, loggedInstitution string) error {

	_, error := usersqueries.GetUserByIDOneInstitution(fmt.Sprintf("%v", teacher), loggedInstitution)
	return error

}

func userTypeVerificationAdding(loggedUser string, loggedInstitution string) (string, bool) {

	userID, _ := usersqueries.GetUserByIDOneInstitution(loggedUser, loggedInstitution)

	if userID.Profile != "Administrador" {
		return "", true
	}
	return userID.Institution, false
}

func verifyIfStudentExist(users primitive.M, loggedInstitution string) string {

	errors := false
	wrongUsers := "Usuarios invalidos o no registrados: \n"
	for user:= range users {
		_, error := usersqueries.GetUserByIDOneInstitution(fmt.Sprintf("%v", user), loggedInstitution)


		if error != nil {
			wrongUsers += fmt.Sprintf("%v", user) + "\n"
			errors = true
		}
	}

	if errors {
		return wrongUsers
	}
	return ""

}

//VerifyIfSubjectExist valida si la asignatura existe
func VerifyIfSubjectExist(subject string, institutionID string) string {

	institutionInfo, found, error := institutionsqueries.GetInstitutionByID(institutionID)

	if error != nil {
		return "Fallo al buscar la institucion"
	}

	if !found {
		return "La institucion no existe"
	}

	_, found = institutionInfo.Subjetcs[subject]

	if !found {

		return "Esta asignatura no existe en la institución "

	}

	return ""

}
