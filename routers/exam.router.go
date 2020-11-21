package routers

import (
	//"fmt"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	//"strconv"

	database "github.com/CJN-Team/examanager-server/database/examqueries"
	generateExam "github.com/CJN-Team/examanager-server/database/generatexamqueries"
	questionsDB "github.com/CJN-Team/examanager-server/database/questionsqueries"
	dbuser "github.com/CJN-Team/examanager-server/database/usersqueries"

	grupDB "github.com/CJN-Team/examanager-server/database/groupqueries"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CreateExam funcion para crear un examen
func CreateExam(w http.ResponseWriter, r *http.Request) {
	var exam models.Exam

	error := json.NewDecoder(r.Body).Decode(&exam)

	user, _ := dbuser.GetUserByIDOneInstitution(IDUser, InstitutionID)

	if user.Profile == "Estudiante" {
		http.Error(w, "La persona no tiene los permisos", 400)
		return
	}

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
	if !exam.MockExam {

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
		status, error = database.UpdateExam(exam, ID, InstitutionID)
	} else {
		id, status, error := generateExam.GenerateMockExam(exam, IDUser, InstitutionID)

		exam.GenerateExam = append(exam.GenerateExam, id)
		if error != nil {
			http.Error(w, "Error al intentar añadir un registro"+error.Error(), 400)
			return
		}

		if status == false {
			http.Error(w, "No se logro añadir un registro"+error.Error(), 400)
			return
		}
		status, error = database.UpdateExam(exam, ID, InstitutionID)
	}

	CleanToken()
	w.WriteHeader(http.StatusCreated)
}

//DeleteExam elimina el examen padre y todos los examenes generados a partir de este.
func DeleteExam(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	user, _ := dbuser.GetUserByIDOneInstitution(IDUser, InstitutionID)
	if user.Profile == "Estudiante" {
		http.Error(w, "La persona no tiene los permisos", 400)
		return
	}

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

//UpdateExam actualiza la nota de un examen generado
func UpdateExam(w http.ResponseWriter, r *http.Request) {
	var exam models.Exam

	error := json.NewDecoder(r.Body).Decode(&exam)

	if error != nil {
		http.Error(w, "Datos Incorrectos"+error.Error(), 400)
		return
	}

	user, _ := dbuser.GetUserByIDOneInstitution(IDUser, InstitutionID)
	if user.Profile == "Estudiante" {
		http.Error(w, "La persona no tiene los permisos", 400)
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
	status, error := database.UpdateExam(exam, id, InstitutionID)

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

//UpdateExamGrade actualiza la nota de un examen generado
func UpdateExamGrade(w http.ResponseWriter, r *http.Request) {
	var exam models.GenerateExam

	error := json.NewDecoder(r.Body).Decode(&exam)

	if error != nil {
		http.Error(w, "Datos Incorrectos"+error.Error(), 400)
		return
	}

	user, _ := dbuser.GetUserByIDOneInstitution(IDUser, InstitutionID)
	if user.Profile == "Estudiante" {
		http.Error(w, "La persona no tiene los permisos"+error.Error(), 400)
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

	user, _ := dbuser.GetUserByIDOneInstitution(IDUser, InstitutionID)
	if user.Profile == "Estudiante" {
		http.Error(w, "La persona no tiene los permisos", 400)
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

//GeneratePDF genera los pdf
func GeneratePDF(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Falta el parametro ID", http.StatusBadRequest)
		return
	}

	user, _ := dbuser.GetUserByIDOneInstitution(IDUser, InstitutionID)
	if user.Profile == "Estudiante" {
		http.Error(w, "La persona no tiene los permisos", 400)
		return
	}

	exam, _, _ := database.GetExamByID(ID, InstitutionID)

	generateExam.CreatePDF(exam, InstitutionID)

	CleanToken()
	w.WriteHeader(http.StatusAccepted)
}

//DownloadPDF descarga el pdf
func DownloadPDF(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	file, error := os.Open("exam-pdf/" + ID + ".pdf")

	if error != nil {
		http.Error(w, "Error al abrir el examen  "+error.Error(), 400)
		return
	}

	_, error = io.Copy(w, file)

	if error != nil {
		http.Error(w, "Error al copiar el examen "+error.Error(), 400)
		return
	}

	CleanToken()
}

//GradeExam califica automaticamente el examen y lo guarda en la base de datos
func GradeExam(w http.ResponseWriter, r *http.Request){
	requestBody := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil{
		http.Error(w, "Debe enviar un request body", http.StatusBadRequest)
		return
	}

	

	examid, found := requestBody["examid"].(string)
	if !found{
		http.Error(w, "Debe especificar el ID del examen", http.StatusBadRequest)
		return
	}

	userAnswers, found := requestBody["questions"].(map[string]interface{})
	if !found{
		http.Error(w, "Debe especificar las preguntas", http.StatusBadRequest)
		return
	}
	

	var generatedExam models.GenerateExam
	generatedExam, found = generateExam.GetGenerateExamByID(examid,InstitutionID)
	if !found{
		http.Error(w, "El examen no existe en esta institucion", http.StatusBadRequest)
		return
	}

	updateString, message := GradeQuestions(generatedExam, userAnswers, InstitutionID)
	if message != ""{
		http.Error(w, "Error al calificar el examen", http.StatusInternalServerError)
		return
	}

	err := generateExam.UpdateExam(examid,updateString)
	if err != nil{
		http.Error(w, "Error al calificar el examen" + err.Error(), http.StatusInternalServerError)
		return
	}

	CleanToken()
	w.WriteHeader(http.StatusAccepted)
}

//GradeQuestions califica automaticamente las respuestas de los estudiantes
func GradeQuestions(generatedExam models.GenerateExam, userAnswers map[string]interface{}, institutionid string)(bson.M, string){
	updateString := bson.M{}
	questionsMap := make(map[string]interface{})
	grade := 0.0
	quantity := 0

	examQuestions := generatedExam.Questions
	for key := range(examQuestions){
		quantity++

		var question models.Question
		question, _, err := questionsDB.GetQuestionByID(key, institutionid)
		if err != nil{
			return updateString, "Error al buscar la pregunta en la base de datos: " + err.Error()
		}
		
		userAnswer := userAnswers[key].([]interface{})
		if question.Category == "Respuesta única" ||  question.Category == "Verdadero o falso"{

			userOption := []string{question.Options[question.Answer[0]]}
			examCorrectOption := []string{question.Options[question.Answer[0]]}

			if userAnswer[0].(int) == question.Answer[0]{
				questionsMap[key] = []interface{}{5.0,userOption,examCorrectOption}
				grade+=5.0

			}else{
				questionsMap[key] = []interface{}{0.0,userOption,examCorrectOption}
			}

		}else if question.Category == "Selección múltiple" {
			goodAnswers := 0
			userOptions := make([]string,len(userAnswer))
			examCorrectOptions := make([]string,len(question.Answer))

			for i, userValue := range userAnswer{
				for j, examValue := range question.Answer{
					if userValue.(int) == examValue{
						goodAnswers++
						examCorrectOptions[j] = question.Options[examValue]
					}
				}
				userOptions[i] = question.Options[userValue.(int)]
			}

			if goodAnswers == len(question.Answer){
				grade+=5.0
				questionsMap[key] = []interface{}{5.0,userOptions,examCorrectOptions}
				
			}else if (goodAnswers < len(question.Answer)) && (goodAnswers > 0){
				grade+=3.0
				questionsMap[key] = []interface{}{3.0,userOptions,examCorrectOptions}

			}else{
				questionsMap[key] = []interface{}{0.0,userOptions,examCorrectOptions}
			}
		}else if question.Category == "Pregunta abierta"{
			questionsMap[key] = []interface{}{0.0,userAnswer,[]string{""}}
		}
	}
	grade /= float64(quantity)
	updateString = bson.M{
		"$set" : bson.M{
			"grade" : grade,
			"question" : questionsMap,
		},
	}

	return updateString, ""
}