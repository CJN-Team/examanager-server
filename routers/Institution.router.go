package routers

import (
	"encoding/json"
	"net/http"

	database "github.com/CJN-Team/examanager-server/database/institutionsqueries"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InstitutionRegistration permite crear una institucion nueva en la base de datos con el modelo de institucion
func InstitutionRegistration(w http.ResponseWriter, r *http.Request) {
	var InstitutionInfo models.Institution
	r.ParseForm()

	err := json.NewDecoder(r.Body).Decode(&InstitutionInfo)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(InstitutionInfo.Name) < 0 {
		http.Error(w, "El nombre es requerido", 400)
		return
	}
	if len(InstitutionInfo.Address) < 0 {
		http.Error(w, "La direccion de la institucion es requerida", 400)
		return
	}
	if len(InstitutionInfo.Phone) < 0 {
		http.Error(w, "El telefono es requerido", 400)
		return
	}
	if len(InstitutionInfo.Type) < 0 {
		http.Error(w, "El tipo de institución es requerido", 400)
		return
	}

	_, found, _ := database.GetInstitutionByName(InstitutionInfo.Name)
	if found {
		http.Error(w, "Ya existe una institución con ese nombre", 400)
		return
	}

	
	if err != nil {
		http.Error(w, "Ha ocurrido un error al crear el documento de usuarios de la institucion "+err.Error(), 400)
		return
	}


	if err != nil {
		http.Error(w, "Ha ocurrido un error al crear el documento de preguntas de la institucion "+err.Error(), 400)
		return
	}



	aux := primitive.M{}
	InstitutionInfo.Subjetcs = aux
	idInstitution, status, err := database.AddInstitution(InstitutionInfo)
	if err != nil {
		http.Error(w, "Ha ocurrido un error al intentar realizar el registro de institucion "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado insertar la institucion nueva ", 400)
		return
	}

	institutionResponse := models.AnswerInstitution{
		InstitutionID: idInstitution,
	}
	CleanToken()
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(institutionResponse)
	w.WriteHeader(http.StatusCreated)
}
