package routers

import (
	"encoding/json"
	"net/http"
	//"strconv"

	database "github.com/CJN-Team/examanager-server/database/questionsqueries"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	if len(question.Thematic) == 0 {
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

	if question.Category == "Pregunta abierta"{
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

	_, found, _ := database.GetQuestionByID(primitive.ObjectID.Hex(question.ID))
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

	var question models.Question

	error := json.NewDecoder(r.Body).Decode(&question)

	if error != nil {
		http.Error(w, "Datos Incorrectos"+error.Error(), 400)
		return
	}

	//status, error := database.UpdateQuestion(question, IDUser)

	if error != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro"+error.Error(), 400)
		return
	}

	/*if status == false {
		http.Error(w, "Ocurrio un error al buscar el registro"+error.Error(), 400)
		return
	}*/

	w.WriteHeader(http.StatusCreated)
}