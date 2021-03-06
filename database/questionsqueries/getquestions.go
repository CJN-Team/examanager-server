package questionsqueries

import (
	"context"
	"log"
	"strconv"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetQuestionByID busca en la base de datos la existencia de una pregunta por el ID
func GetQuestionByID(QuestionID string, institution string) (models.Question, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Questions")

	var questionInfo models.Question
	id := QuestionID
	err := col.FindOne(ctx, bson.M{"_id": id, "institution": institution}).Decode(&questionInfo)

	if err != nil {
		return questionInfo, false, err
	}
	return questionInfo, true, nil
}

//GetAllQuestions retorna todas las preguntas
func GetAllQuestions(category string, category2 string, specific int, page int64, institution string) ([]*models.Question, bool) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := dbConnection.MongoConexion.Database("examanager_db")
	col := db.Collection("Questions")

	var result []*models.Question
	condicion := bson.M{"institution": institution}

	if category2 == "" {
		switch specific {
		case 1:
			condicion = bson.M{
				"topic":       category,
				"institution": institution,
			}
		case 2:
			condicion = bson.M{
				"subject":     category,
				"institution": institution,
			}
		case 3:
			condicion = bson.M{
				"category":    category,
				"institution": institution,
			}
		case 4:
			difficulty, error := strconv.Atoi(category)
			if error != nil {
				return result, false
			}
			aux := int(difficulty)
			condicion = bson.M{
				"difficulty":  aux,
				"institution": institution,
			}
		}
	} else {
		switch specific {
		case 1:
			condicion = bson.M{
				"topic":       category,
				"subject":     category2,
				"institution": institution,
			}
		case 2:
			condicion = bson.M{
				"topic":       category,
				"category":    category2,
				"institution": institution,
			}
		case 3:
			difficulty, error := strconv.Atoi(category2)
			if error != nil {
				return result, false
			}
			aux := int(difficulty)
			condicion = bson.M{
				"topic":       category,
				"difficulty":  aux,
				"institution": institution,
			}
		case 4:
			condicion = bson.M{
				"subject":     category,
				"category":    category2,
				"institution": institution,
			}
		case 5:
			difficulty, error := strconv.Atoi(category2)
			if error != nil {
				return result, false
			}
			aux := int(difficulty)
			condicion = bson.M{
				"subject":     category,
				"difficulty":  aux,
				"institution": institution,
			}
		case 6:
			difficulty, error := strconv.Atoi(category2)
			if error != nil {
				return result, false
			}
			aux := int(difficulty)
			condicion = bson.M{
				"category":    category,
				"difficulty":  aux,
				"institution": institution,
			}
		}
	}

	searchOptions := options.Find()
	if page != -1 {
		searchOptions.SetLimit(20)
		searchOptions.SetSort(bson.D{{Key: "subject", Value: -1}})
		searchOptions.SetSkip((page - 1) * 20)
	}
	pointer, error := col.Find(contex, condicion, searchOptions)

	if error != nil {
		log.Fatal(error.Error())
		return result, false
	}

	for pointer.Next(context.TODO()) {
		var register models.Question
		error := pointer.Decode(&register)

		if error != nil {
			return result, false
		}
		result = append(result, &register)
	}

	return result, true

}

//GetQuestionxInstitution busca en la base de datos el documento que relaciona a las preguntas con la institucion
