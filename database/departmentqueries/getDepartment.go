package departmentqueries

import (
	"context"
	"log"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllDepartments se encarga de traer de base de datos todos los departamentos almacenados
func GetAllDepartments(page int64, institution string) ([]*models.Department, bool) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("departments")

	var result []*models.Department

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
		var register models.Department
		error := pointer.Decode(&register)

		if error != nil {
			return result, false
		}
		result = append(result, &register)
	}

	return result, true
}

//GetDepartmentByID se encarga de buscar en la base de datos el departamento que posee la ID asignada
func GetDepartmentByID(ID string, institution string) (models.Department, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("departments")

	var result models.Department

	departmentID,_:= primitive.ObjectIDFromHex(ID)

	condicion := bson.M{"_id": departmentID, "institution": institution}

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	if error != nil {
		return result, error
	}

	return result, nil
}
