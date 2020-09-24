package routers

import (
	"encoding/json"
	"net/http"
	database "github.com/CJN-Team/examanager-server/database/institutionqueries"
	"github.com/CJN-Team/examanager-server/models"

)

//CreateSubject permite crear una institucion nueva en la base de datos con el modelo de institucion
func CreateSubject(w http.ResponseWriter, r *http.Request) {
	
	var SubjectInfo models.Subject
	err := json.NewDecoder(r.Body).Decode(&SubjectInfo)
	
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(SubjectInfo.Name) < 0 {
		http.Error(w, "El nombre de asignatura es requerido", 400)
		return
	}
	if Profile != "admin" {
		http.Error(w, "Esta opción es válida únicamente para administradores", 403)
		return
	}

	institutionInfo,found,err := database.GetInstitutionByID(InstitutionID)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al buscar el documento de la institucion "+err.Error(), 400)
		return
	}	
	if !found {
		http.Error(w, "La institucion no existe", 400)
		return
	}	
	status, err := database.AddSubject(institutionInfo,SubjectInfo)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al intentar añadir la asignatura "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar la asignatura nueva ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
//DeleteSubject le permite a un administrador de una institucion eliminar una asignatura
func DeleteSubject(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	SubjectName := r.Form.Get("name")	
	if SubjectName == "" {
		http.Error(w, "Error en los datos recibidos ", 400)
		return
	}
	if len(SubjectName) < 0 {
		http.Error(w, "El nombre de la asignatura a eliminar es requerido", 400)
		return
	}
	if Profile != "admin" {
		http.Error(w, "Esta opción es válida únicamente para administradores", 403)
		return
	}

	institutionInfo,found,err := database.GetInstitutionByID(InstitutionID)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al buscar el documento de la institucion "+err.Error(), 400)
		return
	}	
	if !found {
		http.Error(w, "La institucion no existe", 400)
		return
	}	

	_,found = institutionInfo.Subjetcs[SubjectName]

	if !found {
		http.Error(w, "Esta asignatura no existe en la institución ", 406)
		return
	}

	status, err := database.DeleteSubject(institutionInfo,SubjectName)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al intentar eliminar la asignatura "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar la asignatura nueva ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}