package usersqueries

import (
	"context"
	"errors"
	"fmt"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

//UpdateUser se encarga de actualizar el usuario registrado
func UpdateUser(user models.User, ID string, loggedUser string,loggedInstitution string) (bool, error) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	userRegisterd := make(map[string]interface{})

	if len(user.Profile) > 0 {
		userRegisterd["profile"] = user.Profile
	}

	if len(user.IDType) > 0 {
		userRegisterd["idType"] = user.IDType
	}

	if len(user.Name) > 0 {
		userRegisterd["name"] = user.Name
	}

	if len(user.Institution) > 0 {
		userRegisterd["institution"] = user.Institution
	}

	if len(user.LastName) > 0 {
		userRegisterd["lastName"] = user.LastName
	}

	if len(user.Email) > 0 {
		usermodel, exist, _ := GetUserByEmail(user.Email)

		if !exist {
			userRegisterd["email"] = user.Email
		} else if ID == usermodel.ID {
			userRegisterd["email"] = user.Email

		} else {
			error := errors.New("el email ya esta en uso")
			return false, error
		}

	}

	userRegisterd["birthDate"] = user.BirthDate

	if len(user.Photo) > 0 {
		userRegisterd["photo"] = user.Photo
	}

	if len(user.Password) > 0 {
		Password, _ := PasswordEncrypt(user.Password)
		userRegisterd["password"] = Password
	}

	updateString := bson.M{
		"$set": userRegisterd,
	}

	if userTypeVerificationdeleting(loggedUser,loggedInstitution) {
		error := errors.New("el usuario no posee los permisos suficientes")
		return false, error
	}

	fmt.Println(ID)

	filter := bson.M{"_id": bson.M{"$eq": ID},"institution":bson.M{"$eq": ID}}

	_, error := coleccion.UpdateOne(contex, filter, updateString)

	if error != nil {
		return false, error
	}

	return true, nil
}

func userTypeVerificationUpdating(loggedUser string,loggedInstitution string) bool {

	userID, _ := GetUserByIDOneInstitution(loggedUser,loggedInstitution)

	if userID.Profile != "Administrador" {
		return true
	}
	return false
}
