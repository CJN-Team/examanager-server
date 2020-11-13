package routers

import (
	"encoding/json"
	"net/http"

	//"strconv"

	"fmt"

	database "github.com/CJN-Team/examanager-server/database/examqueries"
	generateExam "github.com/CJN-Team/examanager-server/database/generateexamqueries"
	
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

	_, found, _ := database.GetExamByName(exam.Name)
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

	exam, _, _ := database.GetExamByID(ID)

	ids, status, error := generateExam.GenerateExam(exam, IDUser, InstitutionID)

	exam.GenerateExam = ids
	fmt.Println("esto", exam.GenerateExam)
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
