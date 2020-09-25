package usersqueries

import (
	"context"
	"errors"
	"fmt"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
)

//AddUser se encarga de añadir a la base de datos un nuevo usuario
func AddUser(u models.User, loggedUser string) (string, bool, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	u.Password, _ = PasswordEncrypt(u.Password)

	fmt.Println(loggedUser)
	if loggedUser != "" {

		if userTypeVerificationAdding(loggedUser) {
			error := errors.New("el usuario no posee los permisos suficientes")
			return "", false, error
		}

	} else {
		u.Profile = "Administrador"
	}
	fmt.Println(loggedUser)

	_, error := coleccion.InsertOne(contex, u)

	if error != nil {
		return "", false, error
	}

	return "", true, nil
}

func userTypeVerificationAdding(loggedUser string) bool {

	userID, _ := GetUserByID(loggedUser)

	if userID.Profile != "Administrador" {
		return true
	}
	return false
}
