package groupqueries

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/database/examqueries"
	"github.com/CJN-Team/examanager-server/database/institutionsqueries"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllGroups se encarga de traer de base de datos todos los grupos almacenados
func GetAllGroups(page int64, institution string) ([]*models.Group, bool) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("groups")

	var result []*models.Group

	condicion := bson.M{
		"institution": institution,
	}

	searchOptions := options.Find()
	searchOptions.SetLimit(20)
	searchOptions.SetSkip((page - 1) * 20)

	pointer, error := coleccion.Find(contex, condicion, searchOptions)

	if error != nil {
		log.Fatal(error.Error())
		return result, false
	}

	for pointer.Next(context.TODO()) {
		var register models.Group
		error := pointer.Decode(&register)

		if error != nil {
			return result, false
		}
		result = append(result, &register)
	}

	return result, true
}

//GetGroupByID se encarga de buscar en la base de datos el grupo que posee la ID asignada
func GetGroupByID(ID string, institution string) (models.Group, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("groups")

	var result models.Group

	condicion := bson.M{"_id": ID, "institution": institution}

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	if error != nil {
		return result, error
	}

	return result, nil
}

//WatchedTopics se encarga de mostrar cuales temas fueron vistos en el grupo
func WatchedTopics(ID string, institution string) (map[string]bool, error) {

	examsmodels, find := examqueries.GetAllExamByGroup(ID, institution, -1)

	result := make(map[string]bool)

	if !find {
		error := errors.New("No hay examenes asociados a el grupo")
		return result, error
	}

	groupModel, error := GetGroupByID(ID, institution)

	if error != nil {
		return result, error
	}

	institutionInfo, found, error := institutionsqueries.GetInstitutionByID(institution)

	if error != nil {
		error := errors.New("No se pudo obtener la institucion asociada")
		return result, error
	}

	if !found {
		error := errors.New("No existe institucion asociada")
		return result, error
	}

	subject, found := institutionInfo.Subjetcs[groupModel.Subject]

	for _, value := range subject.(primitive.A) {
		fmt.Println(value)
		result[value.(string)] = false
	}
	for _, value := range examsmodels {

		result[value.TopicQuestion] = true
	}

	return result, nil
}

