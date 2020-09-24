package userqueries

import "golang.org/x/crypto/bcrypt"

//PasswordEncrypt se encarga de agregar una capa de encryptacion a nuestra contraseña por medio de la libreria Bcrypt
func PasswordEncrypt(password string) (string, error) {
	passwordLevel := 8

	bytes, error := bcrypt.GenerateFromPassword([]byte(password), passwordLevel)

	return string(bytes), error
}
