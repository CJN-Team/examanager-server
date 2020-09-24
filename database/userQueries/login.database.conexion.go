package userqueries

import (
	"github.com/CJN-Team/examanager-server/models"
	"golang.org/x/crypto/bcrypt"
)

//UserLogin se encarga de realizar del usuario por medio de un email y comparacion de la contraseña ingresada y la almacenada
func UserLogin(email string, password string) (models.User, bool) {
	user, find, _ := GetUserByEmail(email)

	if find == false {
		return user, find
	}

	passwordUsed := []byte(password)
	passwordSaved := []byte(user.Password)

	error := bcrypt.CompareHashAndPassword(passwordSaved, passwordUsed)

	if error != nil {
		return user, false
	}

	return user, true
}
