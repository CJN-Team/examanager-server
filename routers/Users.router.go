package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
)

//CreateUser funcion para crear un usuario en la base de datos
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var t models.User

	error := json.NewDecoder(r.Body).Decode(&t)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	}

	//Validaciones de los datos a registrar

	if len(t.ID) == 0 {
		http.Error(w, "La ID es requerida", 400)
		return
	}
	if len(t.Profile) == 0 {
		http.Error(w, "El Perfil es requerido", 400)
		return
	}
	if len(t.LastName) == 0 {
		http.Error(w, "El apellido es requerido", 400)
		return
	}
	if len(t.Name) == 0 {
		http.Error(w, "El nombre es requerido", 400)
		return
	}

	_, found, _ := database.GetUserByEmail(t.Email)

	if found {
		http.Error(w, "El usuario ya existe", 400)
		return
	}

	_, status, error := database.AddUser(t)

	if error != nil {
		http.Error(w, "Error al intentar añadir un registro"+error.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro añadir un registro"+error.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

/*
//ReadUser funcion para la lectura de un usuario presente en la base de datos
func ReadUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	error := json.NewDecoder(r.Body).Decode(&user)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	}

	user, found, _ := database.GetUserByID(user.ID)

	if error != nil {
		http.Error(w, "El usuario no existe", 400)
		return
	}

}
*/
