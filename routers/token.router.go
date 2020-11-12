package routers

import (
	"errors"
	"fmt"
	"strings"

	database "github.com/CJN-Team/examanager-server/database/usersqueries"
	"github.com/CJN-Team/examanager-server/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Email es el email obtenido del modelo y que sera usado en los endpoints
var Email string

//IDUser es el id del usuario obtenido del modelo
var IDUser string

//Profile es el rol de usuario obtenido del modelo
var Profile string

//InstitutionID es el ID de referencia de la Institucion a la que pertenece el usuario
var InstitutionID string

//GetToken permite extraer los valores que contiene el token
func GetToken(token string) (*models.Claim, bool, string, string, string, error) {
	password := []byte("YouShallNotPasssssss")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), string(""), string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tokens, error := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return password, nil
	})

	if error == nil {
		_, found, _ := database.GetUserByEmail(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID
			Profile = claims.Profile
			InstitutionID = claims.Institution
		}
		fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Email, IDUser, Profile, InstitutionID)
		return claims, found, IDUser, Profile, InstitutionID, nil
	}
	fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", Email, IDUser, Profile, InstitutionID)

	if !tokens.Valid {
		return claims, false, string(""), string(""), string(""), errors.New("token invalido")
	}
	return claims, false, string(""), string(""), string(""), error
}

//CleanToken Limpia las variables globales
func CleanToken() {
	Email = ""
	IDUser = ""
	Profile = ""
	InstitutionID = ""
}
