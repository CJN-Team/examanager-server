package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"
	database "github.com/CJN-Team/examanager-server/database/usersqueries"
	"github.com/CJN-Team/examanager-server/models"
)

//CreateUser funcion para crear un usuario en la base de datos
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ID de Usuario :" + IDUser)
	fmt.Println("ID de Institucion :" +InstitutionID)
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

	_, status, error := database.AddUser(t, IDUser)

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
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ID de Usuario :" + IDUser)
	fmt.Println("ID de Institucion :" +InstitutionID)
	var user models.User

	error := json.NewDecoder(r.Body).Decode(&user)

	if error != nil {
		http.Error(w, "Datos Incorrectos"+error.Error(), 400)
		return
	}

	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "Debe enviar el perfil a buscar", http.StatusBadRequest)
		return
	}

	if IDUser == "" {
		http.Error(w, "Debes estar logueado", http.StatusBadRequest)
		return
	}
	status, error := database.UpdateUser(user, id,IDUser)

	if error != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro"+error.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Ocurrio un error al buscar el registro"+error.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

//GetAllUsersRouter permite tomar todos los usuarios de una categoria
func GetAllUsersRouter(w http.ResponseWriter, r *http.Request) {

	profile := r.URL.Query().Get("profile")

	if len(profile) < 1 {
		http.Error(w, "Debe enviar el perfil a buscar", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	page, error := strconv.Atoi(r.URL.Query().Get("page"))

	if error != nil {
		http.Error(w, "Pagina debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	pageAux := int64(page)

	result, correct := database.GetAllUsers(profile, pageAux)

	if correct == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)

}

//DeleteUserRouter elimina el usuario seleccionado
func DeleteUserRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ID de Usuario :" + IDUser)
	fmt.Println("ID de Institucion :" +InstitutionID)
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	error := database.DeleteUser(ID, IDUser)

	if error != nil {
		http.Error(w, "Ocurrio un error al intentar borrar un usuario"+error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
