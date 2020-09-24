package routers

import (
	"encoding/json"
	"net/http"
	database "github.com/CJN-Team/examanager-server/database/institutionqueries"
	"github.com/CJN-Team/examanager-server/models"
)

/*
	CreateSubject permite crear una institucion nueva en la base de datos con el modelo de institucion,
	verificando solamente los datos principales de la institución para poder crearla.
*/
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
		http.Error(w, "Esta opción es válida únicamente para administradores", 400)
		return
	}
	/*if len(SubjectInfo.TopicsList) < 0 {
		http.Error(w, "Las tematicas de la asignatura deben ser validas", 400)
		return
	}*/

	institutionInfo,found,err := database.GetInstitutionByID(InstitutionID)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al buscar el documento de la institucion "+err.Error(), 400)
		return
	}	
	if !found {
		http.Error(w, "La institucion no existe", 400)
		return
	}	
	status, err := database.AddSubject(SubjectInfo,institutionInfo)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al intentar realizar el registro de institucion "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar la asignatura nueva ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
