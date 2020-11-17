package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	//"strconv"

	database "github.com/CJN-Team/examanager-server/database/examqueries"
	generateExam "github.com/CJN-Team/examanager-server/database/generatexamqueries"

	grupDB "github.com/CJN-Team/examanager-server/database/groupqueries"
	"github.com/CJN-Team/examanager-server/models"
)

//CreateExam funcion para crear un examen
func CreateExam(w http.ResponseWriter, r *http.Request) {
	var exam models.Exam

	error := json.NewDecoder(r.Body).Decode(&exam)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	}

	//Validaciones de los datos a registrar

	_, err := grupDB.GetGroupByID(exam.GroupID, InstitutionID)
	if err != nil {
		http.Error(w, "El grupo no existe", 400)
		return
	}

	very := grupDB.VerifyIfSubjectExist(exam.SubjectID, InstitutionID)
	if very != "" {
		http.Error(w, very, 400)
		return
	}

	if len(exam.Name) == 0 {
		http.Error(w, "El nobre del examen es necesario", 400)
		return
	}
	if len(exam.Institution) == 0 {
		http.Error(w, "la institucion es necesaria", 400)
		return
	}
	if exam.State == true {

	} else {
		exam.State = false
	}

	if len(exam.Difficulty) == 0 {
		http.Error(w, "La dificultad es necesaria", 400)
		return
	}

	if len(exam.TopicQuestion) == 0 {
		http.Error(w, "La tematica es requerida", 400)
		return
	}

	_, found, _ := database.GetExamByName(exam.Name, exam.GroupID, InstitutionID)
	if found {
		http.Error(w, "Ya existe un examen con ese nombre", 400)
		return
	}

	_, status, error := database.AddExam(exam)

	if error != nil {
		http.Error(w, "Error al intentar añadir un registro"+error.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro añadir un registro"+error.Error(), 400)
		return
	}
	CleanToken()
	w.WriteHeader(http.StatusCreated)
}

//CreateGenerateExam funcion para crear un examen
func CreateGenerateExam(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Falta el parametro ID", http.StatusBadRequest)
		return
	}

	exam, _, _ := database.GetExamByID(ID, InstitutionID)

	if len(exam.GenerateExam) != 0 {
		http.Error(w, "ya se han generado los examenes de este modelo", 400)
		return
	}

	ids, status, error := generateExam.GenerateExam(exam, IDUser, InstitutionID)

	exam.GenerateExam = ids
	if error != nil {
		http.Error(w, "Error al intentar añadir un registro"+error.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro añadir un registro"+error.Error(), 400)
		return
	}
	status, error = database.UpdateExam(exam, ID)
	CleanToken()
	w.WriteHeader(http.StatusCreated)
}

//DeleteExam elimina el examen padre y todos los examenes generados a partir de este.
func DeleteExam(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Falta el parametro ID", http.StatusBadRequest)
		return
	}
	_, err := database.DeleteExam(ID)
	if err != nil {
		http.Error(w, "Error al eliminar el examen: "+err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = database.DeleteGeneratedExams(ID)
	if err != nil {
		http.Error(w, "Error al eliminar el examen: "+err.Error(), http.StatusInternalServerError)
		return
	}

	CleanToken()
	w.WriteHeader(http.StatusAccepted)
}

//UpdateExamGrade actualiza la nota de un examen generado
func UpdateExamGrade(w http.ResponseWriter, r *http.Request) {
	var exam models.GenerateExam

	error := json.NewDecoder(r.Body).Decode(&exam)

	if error != nil {
		http.Error(w, "Datos Incorrectos"+error.Error(), 400)
		return
	}

	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "Debe enviar el examen a buscar", http.StatusBadRequest)
		return
	}

	if IDUser == "" {
		http.Error(w, "Debes estar logueado", http.StatusBadRequest)
		return
	}
	status, error := database.UpdateGenerateExam(exam, id)

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

//GetAllExams permite tomar todos los examenes de un grupo
func GetAllExams(w http.ResponseWriter, r *http.Request) {

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

	groupID := r.URL.Query().Get("groupid")

	if len(groupID) < 1 {
		http.Error(w, "Falta el parametro groupID", http.StatusBadRequest)
		return
	}

	result, correct := database.GetAllExamByGroup(groupID, InstitutionID, pageAux)

	if correct == false {
		http.Error(w, "Error al leer los grupos", http.StatusBadRequest)
		return
	}
	CleanToken()
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)

}

//GetGenerateExam permite tomar todos los examenes de un grupo
func GetGenerateExam(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Falta el parametro id", http.StatusBadRequest)
		return
	}

	result, correct := generateExam.GetGenerateExamByID(ID, InstitutionID)

	if correct == false {
		http.Error(w, "Error al buscar el examen", http.StatusBadRequest)
		return
	}
	CleanToken()
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)
}
