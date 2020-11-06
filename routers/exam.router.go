package routers

import (
	"encoding/json"
	"net/http"

	//"strconv"

	//"fmt"

	database "github.com/CJN-Team/examanager-server/database/examqueries"
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

	_, err := grupDB.GetGroupByID(exam.GroupID)
	if err != nil {
		http.Error(w, "El grupo no existe", 400)
		return
	}

	if len(exam.Name) == 0 {
		http.Error(w, "El nobre del examen es necesario", 400)
		return
	}

	if exam.NumberQuestions == 0 {
		http.Error(w, "El numero de preguntas es necesario", 400)
		return
	}

	if exam.Difficulty == 0 {
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

	w.WriteHeader(http.StatusCreated)
}
