package routers

import (
	"encoding/json"
	"net/http"
	"time"

	database "github.com/CJN-Team/examanager-server/database/userQueries"
	"github.com/CJN-Team/examanager-server/jwt"
	"github.com/CJN-Team/examanager-server/models"
)

//Login se encarga de realizar la funcion de iniciar sesion de los usuarios
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var user models.User

	error := json.NewDecoder(r.Body).Decode(&user)

	if error != nil {
		http.Error(w, "Usuario o contraseña invalidos"+error.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El Email es requerido", 400)
		return
	}

	userFound, exist := database.UserLogin(user.Email, user.Password)

	if exist == false {
		http.Error(w, "Usuario o contraseña invalidos", 400)
		return
	}

	jwtKey, error := jwt.GenerateJWT(userFound)

	if error != nil {
		http.Error(w, "Error generando el token"+error.Error(), 400)
		return
	}

	AnswerLogin := models.AnswerLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(AnswerLogin)

	//Como generar una cockie

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}


