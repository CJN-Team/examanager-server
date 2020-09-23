package routers

import (
	"errors"
	"strings"

	"github.com/CJN-Team/examanager-server/models"
	database "github.com/CJN-Team/examanager-server/database/userQueries"
	jwt "github.com/dgrijalva/jwt-go"
)


//Email es el email obtenido del modelo y que sera usado en los endpoints
var Email string

//IDUser es el id del usuario obtenido del modelo
var IDUser string

//GetToken permite extraer los valores que contiene el token
func GetToken(token string) (*models.Claim,bool,string, error){
	password := []byte ("YouShallNotPasssssss")
	claims := &models.Claim {}

	splitToken := strings.Split(token,"Bearer")
	if len(splitToken) != 2 {
		return claims,false, string(""),errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tokens, error := jwt.ParseWithClaims(token,claims, func(token *jwt.Token)(interface{},error){
		return password,nil
	})

	if error == nil{
		_,found, _ := database.GetUserByEmail(claims.Email)
		if found == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims,found,IDUser,nil
	}
	if !tokens.Valid{
		return claims,false,string(""),errors.New("token invalido")
	}
	return claims, false, string(""),error
}