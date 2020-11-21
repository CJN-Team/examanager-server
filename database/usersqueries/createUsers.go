package usersqueries

import (
	"context"
	"errors"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
)

//AddUser se encarga de añadir a la base de datos un nuevo usuario
func AddUser(u models.User, loggedUser string, loggedInstitution string) (string, bool, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	u.Password, _ = PasswordEncrypt(u.Password)

	_, found, _ := GetUserByEmail(u.Email)

	if found {
		error := errors.New("El usuario ya existe")
		return "", false, error
	}

	_, error := GetUserByIDAllInstitutions(u.ID)

	if error == nil {
		error := errors.New("El usuario ya existe")
		return "", false, error
	}

	if loggedUser != "" {
		institution, admin := userTypeVerificationAdding(loggedUser, loggedInstitution)
		if admin {
			error := errors.New("el usuario no posee los permisos suficientes")
			return "", false, error
		}
		u.Institution = institution
	} else {
		u.Profile = "Administrador"
	}

	_, error = coleccion.InsertOne(contex, u)

	if error != nil {
		return "", false, error
	}

	return "", true, nil
}

func userTypeVerificationAdding(loggedUser string, loggedInstitution string) (string, bool) {

	userID, _ := GetUserByIDOneInstitution(loggedUser, loggedInstitution)

	if userID.Profile != "Administrador" {
		return "", true
	}
	return userID.Institution, false
}
