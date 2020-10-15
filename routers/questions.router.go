package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	//"strconv"

	//"fmt"

	database "github.com/CJN-Team/examanager-server/database/questionsqueries"
	"github.com/CJN-Team/examanager-server/models"
)

//CreateQuestion funcion para crear un usuario en la base de datos
func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var question models.Question

	error := json.NewDecoder(r.Body).Decode(&question)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	}

	//Validaciones de los datos a registrar

	if len(question.ID) == 0 {
		http.Error(w, "La ID es requerida", 400)
		return
	}
	if len(question.Topic) == 0 {
		http.Error(w, "La tematica es requerida", 400)
		return
	}

	if len(question.Subject) == 0 {
		http.Error(w, "La asignatura es requerida", 400)
		return
	}

	if len(question.Pregunta) == 0 {
		http.Error(w, "La pregunta es requerida", 400)
		return
	}

	if len(question.Category) == 0 {
		http.Error(w, "La categoria es requerida", 400)
		return
	}

	if question.Category == "Pregunta abierta" {
		if len(question.Options) == 0 {
			http.Error(w, "Las opciones requerida", 400)
			return
		}
	}

	if len(question.Answer) == 0 {
		http.Error(w, "La ID es requerida", 400)
		return
	}

	if question.Difficulty == 0 {
		http.Error(w, "La ID es requerida", 400)
		return
	}

	_, found, _ := database.GetQuestionByID(question.ID)
	if found {
		http.Error(w, "Ya existe una pregunta con ese ID", 400)
		return
	}

	_, status, error := database.AddQuestion(question)

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

//UpdateQuestion se encarga de la actualizacion de la pregunta
func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hoola ")
	var question models.Question

	error := json.NewDecoder(r.Body).Decode(&question)

	if error != nil {
		http.Error(w, "Datos Incorrectos"+error.Error(), 400)
		return
	}

	id := r.URL.Query().Get("id")
	fmt.Println(id)

	if len(id) < 1 {
		http.Error(w, "Debe enviar la pregunta a buscar", http.StatusBadRequest)
		return
	}

	status, error := database.UpdateQuestion(question, id)

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

//GetAllQuestionsRouter se encarga de devollver todas las preguntas
func GetAllQuestionsRouter(w http.ResponseWriter, r *http.Request) {

	page, error := strconv.Atoi(r.URL.Query().Get("page"))

	if error != nil {
		http.Error(w, "Pagina debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	pageAux := int64(page)

	category := r.URL.Query().Get("category")
	category2 := r.URL.Query().Get("category2")
	specific, error := strconv.Atoi(r.URL.Query().Get("specific"))

	if error != nil {
		http.Error(w, "specific debe estar definido", http.StatusBadRequest)
		return
	}
	specific = int(specific)

	result, correct := database.GetAllQuestions(category, category2, specific, pageAux)

	if correct == false {
		http.Error(w, "Error al leer las preguntas", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(result)
}

//DeleteQuestionsRouter elimina el usuario seleccionado
func DeleteQuestionsRouter(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	error := database.DeleteQuestion(ID)

	if error != nil {
		http.Error(w, "Ocurrio un error al intentar borrar un usuario"+error.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
