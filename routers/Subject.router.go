package routers

import (
	"encoding/json"
	"net/http"
	database "github.com/CJN-Team/examanager-server/database/institutionsqueries"
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
//UpdateSubject permite editar el nombre de una asignatura y tus tematicas
func UpdateSubject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err!=nil{
		http.Error(w, "Ha ocurrido un error al obtener la peticion "+err.Error(), 400)
		return
	}

	params:=r.Form

	var SubjectInfo models.Subject
	err = json.NewDecoder(r.Body).Decode(&SubjectInfo)

	if err!=nil{
		http.Error(w, "Ha ocurrido un error al obtener la peticion "+err.Error(), 400)
		return
	}

	option, found := params["option"]
	if !found {
		http.Error(w, "Debe enviar una opcion de actualizacion "+err.Error(), 405)
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
	subjectName, found := params["name"]
	if !found{
		http.Error(w, "Debe especificar el nombre de una asignatura a modificar ", 400)
		return
	}
	subject,found := institutionInfo.Subjetcs[subjectName[0]]
	if !found {
		http.Error(w, "Esta asignatura no existe en la institución ", 400)
		return
	}
	if option[0] == "1"{
		response,status, err := UpdateSubjectName(subjectName[0], subject,SubjectInfo,institutionInfo)
		if err != nil {
			http.Error(w, "Ha ocurrido un error al modificar el nombre de la asignatura"+err.Error(), 400)
			return
		}
		if !status {
			http.Error(w, "No se ha logrado modificar el nombre de la asignatura: " + response, 400)
			return
		}
	}else if option[0] == "2"{
		response,status, err := UpdateSubjectTopics(subjectName[0],SubjectInfo,institutionInfo)
		if err != nil {
			http.Error(w, "Ha ocurrido un error al modificar el nombre de la asignatura"+err.Error(), 400)
			return
		}
		if !status {
			http.Error(w, "No se ha logrado modificar el nombre de la asignatura: " + response, 400)
			return
		}
	}else{
		http.Error(w, "La opcion elegida no es valida "+err.Error(), 405)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
//UpdateSubjectName sdasd
func UpdateSubjectName(subjectName string, subject interface{},
		SubjectInfo models.Subject, institutionInfo models.Institution)(string,bool,error){

	institutionInfo.Subjetcs[SubjectInfo.Name] = subject
	delete(institutionInfo.Subjetcs, subjectName)
	status, err := database.DeleteSubject(institutionInfo, subjectName)
	if err != nil{
		return "Error al eliminar la asignatura anterior ", false, err
	}
	if !status{
		return "Error al eliminar la asignatura anterior ", false, nil
	}
	status, err = database.AddSubject(institutionInfo)
	if err != nil{
		return "Error al crear la asignatura nueva ", false, err
	}
	if !status{
		return "Error al crear la asignatura nueva ", false, nil
	}
	return string(""), true, nil
}
//UpdateSubjectTopics permite 
func UpdateSubjectTopics(subjectName string,SubjectInfo models.Subject, institutionInfo models.Institution)(string,bool,error){
	delete(institutionInfo.Subjetcs, subjectName)
	institutionInfo.Subjetcs[SubjectInfo.Name] = SubjectInfo.Topics
	status, err := database.DeleteSubject(institutionInfo, subjectName)
	if err != nil{
		return "Error al eliminar la asignatura anterior ", false, err
	}
	if !status{
		return "Error al eliminar la asignatura anterior ", false, nil
	}
	status, err = database.AddSubject(institutionInfo)
	if err != nil{
		return "Error al crear la asignatura nueva ", false, err
	}
	if !status{
		return "Error al crear la asignatura nueva ", false, nil
	}
	return string(""), true, nil
}