package jwt

import (
	"time"

	"github.com/CJN-Team/examanager-server/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GenerateJWT Genera el Json web token que utilizara al momento de enviar los datos
func GenerateJWT(user models.User) (string, error) {

	password := []byte("YouShallNotPasssssss")

	payload := jwt.MapClaims{
		"email":       user.Email,
		"profile":     user.Profile,
		"idType":      user.IDType,
		"name":        user.Name,
		"lastName":    user.LastName,
		"birthDate":   user.Email,
		"_id":         user.ID.Hex(),
		"expiration":  time.Now().Add(time.Hour * 24).Unix(),
		"institution": user.Institution,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, error := token.SignedString(password)

	if error != nil {
		return tokenString, error
	}

	return tokenString, nil
}
