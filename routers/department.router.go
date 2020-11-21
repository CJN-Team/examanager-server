package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	database "github.com/CJN-Team/examanager-server/database/departmentqueries"
	"github.com/CJN-Team/examanager-server/models"
)

//CreateDepartment funcion para crear un departamento en la base de datos
func CreateDepartment(w http.ResponseWriter, r *http.Request) {
	var department models.Department

	department.Institution = InstitutionID

	error := json.NewDecoder(r.Body).Decode(&department)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	}

	//Validaciones de los datos a registrar

	if len(department.Name) == 0 {
		http.Error(w, "El Nombre del departamento es requerido", 400)
		return
	}

	if len(department.Teachers) == 0 {
		department.Teachers = []string{}
	}

	_, status, error := database.AddDepartment(department, IDUser, InstitutionID)

	if error != nil {
		http.Error(w, "Error al intentar añadir un registro: "+error.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro añadir un registro: "+error.Error(), 400)
		return
	}
	CleanToken()
	w.WriteHeader(http.StatusCreated)
}

//ReadDepartment funcion para la lectura de un departamento especifico presente en la base de datos
func ReadDepartment(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Falta el parametro ID", http.StatusBadRequest)
		return
	}

	user, error := database.GetDepartmentByID(ID, InstitutionID)

	if error != nil {
		http.Error(w, "Ocurrio un error al buscar el registro"+error.Error(), 400)
		return
	}
	CleanToken()
	w.Header().Set("context-type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}

//GetAllDepartments permite tomar todos los departamentos de una categoria
func GetAllDepartments(w http.ResponseWriter, r *http.Request) {

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

	result, correct := database.GetAllDepartments(pageAux, InstitutionID)

	if correct == false {
		http.Error(w, "Error al leer los departamentos", http.StatusBadRequest)
		return
	}
	CleanToken()
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)

}

//DeleteDepartment elimina el departamento seleccionado
func DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	error := database.DeleteDepartment(ID, IDUser,InstitutionID)

	if error != nil {
		http.Error(w, "Ocurrio un error al intentar borrar un departamento"+error.Error(), http.StatusBadRequest)
		return
	}
	CleanToken()
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

//UpdateDepartment se encarga de la actualizacion del departamento seleccionado
func UpdateDepartment(w http.ResponseWriter, r *http.Request) {

	var department models.Department

	error := json.NewDecoder(r.Body).Decode(&department)

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
	status, error := database.UpdateDepartment(department, id, IDUser,InstitutionID)

	if error != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro"+error.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Ocurrio un error al buscar el registro"+error.Error(), 400)
		return
	}
	CleanToken()
	w.WriteHeader(http.StatusCreated)
}
