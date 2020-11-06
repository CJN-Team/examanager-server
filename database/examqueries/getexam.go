package examqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
)

//GetExamByName verifica si el usuario ya se encuentra en la base de datos por medio de el correo
func GetExamByName(name string) (models.Exam, bool, string) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")
	coleccion := database.Collection("Exam")

	condicion := bson.M{"name": name}

	var result models.Exam

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	ID := result.ID

	if error != nil {
		return result, false, ID.Hex()
	}

	return result, true, ID.Hex()
}
