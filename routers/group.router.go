package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	database "github.com/CJN-Team/examanager-server/database/groupqueries"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateGroup funcion para crear un grupo en la base de datos
func CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group models.Group

	error := json.NewDecoder(r.Body).Decode(&group)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	}

	//Validaciones de los datos a registrar

	if len(group.ID) == 0 {
		http.Error(w, "La ID del grupo es requerida", 400)
		return
	}

	if len(group.Name) == 0 {
		http.Error(w, "El Nombre del grupo es requerido", 400)
		return
	}

	if len(group.Teacher) == 0 {
		http.Error(w, "El Profesor del grupo es requerido", 400)
		return
	}

	if len(group.StudentsList) == 0 {
		group.StudentsList = primitive.M{}
	}

	_, status, error := database.AddGroup(group, IDUser)

	if error != nil {
		http.Error(w, "Error al intentar añadir un registro: "+error.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro añadir un registro: "+error.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

//ReadGroup funcion para la lectura de un grupo especifico presente en la base de datos
func ReadGroup(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Falta el parametro ID", http.StatusBadRequest)
		return
	}

	user, error := database.GetGroupByID(ID)

	if error != nil {
		http.Error(w, "Ocurrio un error al buscar el registro"+error.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}

//GetAllGroups permite tomar todos los grupos de una categoria
func GetAllGroups(w http.ResponseWriter, r *http.Request) {

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

	result, correct := database.GetAllGroups(pageAux)

	if correct == false {
		http.Error(w, "Error al leer los grupos", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)

}

//DeleteGroup elimina el grupo seleccionado
func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	fmt.Println(ID, IDUser)

	error := database.DeleteGroup(ID, IDUser)

	if error != nil {
		http.Error(w, "Ocurrio un error al intentar borrar un grupo"+error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

//UpdateGroup se encarga de la actualizacion del grupo seleccionado
func UpdateGroup(w http.ResponseWriter, r *http.Request) {

	var group models.Group

	error := json.NewDecoder(r.Body).Decode(&group)

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
	status, error := database.UpdateGroup(group, id, IDUser)

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