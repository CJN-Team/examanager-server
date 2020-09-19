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

//ReadUser funcion para la lectura de un usuario presente en la base de datos
func ReadUser(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Falta el parametro ID", http.StatusBadRequest)
		return
	}

	user, error := database.GetUserByID(ID)

	if error != nil {
		http.Error(w, "Ocurrio un error al buscar el registro"+error.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}

//UpdateUser se encarga de la actualizacion del usuario seleccionado
func UpdateUser(w http.ResponseWriter, r *http.Request){

	var user models.User

	error := json.NewDecoder(r.Body).Decode(&user)

	if error != nil{
		http.Error(w, "Datos Incorrectos"+error.Error(),400)
		return
	}

	status,error := database.UpdateUser(user,IDUser)

	if error != nil{
		http.Error(w,"Ocurrio un error al intentar modificar el registro"+error.Error(),400)
		return
	}

	if status == false {
		http.Error(w, "Ocurrio un error al buscar el registro"+error.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
