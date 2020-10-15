package questionsqueries

import (
	"context"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
)

//AddQuestion inserta en la base de datos el modelo de la pregunta
func AddQuestion(questionModel models.Question) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Questions")

	_, err := col.InsertOne(ctx, questionModel)

	id := questionModel.ID
	if err != nil {
		return id, false, err
	}

	return id, true, nil
}
