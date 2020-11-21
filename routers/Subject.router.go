package routers

import (
	"encoding/json"
	"net/http"

	database "github.com/CJN-Team/examanager-server/database/institutionsqueries"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
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
	if Profile != "Administrador" {
		http.Error(w, "Esta opción es válida únicamente para administradores", 403)
		return
	}

	institutionInfo, found, err := database.GetInstitutionByID(InstitutionID)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al buscar el documento de la institucion "+err.Error(), 400)
		return
	}
	if !found {
		http.Error(w, "La institucion no existe", 400)
		return
	}
	institutionInfo.Subjetcs[SubjectInfo.Name] = SubjectInfo.Topics
	status, err := database.AddSubject(institutionInfo)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al intentar añadir la asignatura "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar la asignatura nueva ", 400)
		return
	}
	CleanToken()
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
	if Profile != "Administrador" {
		http.Error(w, "Esta opción es válida únicamente para administradores", 403)
		return
	}

	institutionInfo, found, err := database.GetInstitutionByID(InstitutionID)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al buscar el documento de la institucion "+err.Error(), 400)
		return
	}
	if !found {
		http.Error(w, "La institucion no existe", 400)
		return
	}

	_, found = institutionInfo.Subjetcs[SubjectName]

	if !found {
		http.Error(w, "Esta asignatura no existe en la institución ", 406)
		return
	}

	status, err := database.DeleteSubject(institutionInfo, SubjectName)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al intentar eliminar la asignatura "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar la asignatura nueva ", 400)
		return
	}
	CleanToken()
	w.WriteHeader(http.StatusCreated)
}

//UpdateSubject permite editar el nombre de una asignatura y tus tematicas
func UpdateSubject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ha ocurrido un error al obtener la peticion "+err.Error(), 400)
		return
	}

	params := r.Form

	var SubjectInfo models.Subject
	err = json.NewDecoder(r.Body).Decode(&SubjectInfo)

	if err != nil {
		http.Error(w, "Ha ocurrido un error al obtener la peticion "+err.Error(), 400)
		return
	}

	institutionInfo, found, err := database.GetInstitutionByID(InstitutionID)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al buscar el documento de la institucion "+err.Error(), 400)
		return
	}
	if !found {
		http.Error(w, "La institucion no existe", 400)
		return
	}
	subjectName, found := params["name"]
	if !found {
		http.Error(w, "Debe especificar el nombre de una asignatura a modificar ", 400)
		return
	}
	_, found = institutionInfo.Subjetcs[subjectName[0]]
	if !found {
		http.Error(w, "Esta asignatura no existe en la institución ", 400)
		return
	}

	response, status, err := UpdateSubjectTopics(subjectName[0], SubjectInfo, institutionInfo)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al modificar la asignatura"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar la asignatura: "+response, 400)
		return
	}
	CleanToken()
	w.WriteHeader(http.StatusCreated)
}

//UpdateSubjectTopics permite
func UpdateSubjectTopics(subjectName string, SubjectInfo models.Subject, institutionInfo models.Institution) (string, bool, error) {
	delete(institutionInfo.Subjetcs, subjectName)
	institutionInfo.Subjetcs[SubjectInfo.Name] = SubjectInfo.Topics
	status, err := database.AddSubject(institutionInfo)
	if err != nil {
		return "Error al crear la asignatura nueva ", false, err
	}
	if !status {
		return "Error al crear la asignatura nueva ", false, nil
	}
	CleanToken()
	return string(""), true, nil
}

//GetSubjects trata de recuperar las asignaturas y las materias de una institucion
func GetSubjects(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	SubjectName := r.Form.Get("name")
	if SubjectName != "" {
		GetSubject(w, r, SubjectName)
		return
	}
	institutionInfo, found, err := database.GetInstitutionByID(InstitutionID)
	if err != nil {
		http.Error(w, "Error al buscar la institucion: "+err.Error(), 400)
		return
	}
	if !found {
		http.Error(w, "El usuario no está asociado a una institución existente ", 400)
		return
	}
	Sujects := institutionInfo.Subjetcs

	CleanToken()

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(Sujects)
	w.WriteHeader(http.StatusCreated)
}

//GetSubject trata de recuperar una asignatura
func GetSubject(w http.ResponseWriter, r *http.Request, SubjectName string) {

	if SubjectName == "" {
		http.Error(w, "Error en los datos recibidos ", 400)
		return
	}
	if len(SubjectName) < 0 {
		http.Error(w, "El nombre de la asignatura a eliminar es requerido", 400)
		return
	}

	institutionInfo, found, err := database.GetInstitutionByID(InstitutionID)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al buscar el documento de la institucion "+err.Error(), 400)
		return
	}
	if !found {
		http.Error(w, "La institucion no existe", 400)
		return
	}

	subject, found := institutionInfo.Subjetcs[SubjectName]
	subjectInfo := bson.M{
		SubjectName: subject,
	}
	if !found {
		http.Error(w, "Esta asignatura no existe en la institución ", 406)
		return
	}
	CleanToken()
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(subjectInfo)
	w.WriteHeader(http.StatusCreated)

}
