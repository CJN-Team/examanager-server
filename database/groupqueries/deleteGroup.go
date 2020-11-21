package groupqueries

import (
	"context"
	"errors"
	"time"

	//examqueries "github.com/CJN-Team/examanager-server/database/examqueries"
	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/database/usersqueries"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeleteGroup se encarga de borrar el grupo seleccionado
func DeleteGroup(ID string, loggedUser string, loggedInstitution string) error {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("groups")

	groupID := ID

	condicion := bson.M{
		"_id":         groupID,
		"institution": loggedInstitution,
	}

	if userTypeVerificationdeleting(loggedUser, loggedInstitution) {
		error := errors.New("el usuario no posee los permisos suficientes")
		return error
	}

	exist, error := coleccion.DeleteOne(contex, condicion)
	if exist.DeletedCount == 0 {
		error := errors.New("el archivo a eliminar no existe")
		return error
	}
	return error
}

func userTypeVerificationdeleting(loggedUser string, loggedInstitution string) bool {

	userID, _ := usersqueries.GetUserByIDOneInstitution(loggedUser, loggedInstitution)

	if userID.Profile != "Administrador" {
		return true
	}
	return false
}

func DeleteExamsOfStudents(groupid string, generatedExamID string, institutionID string)(error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")
	col := database.Collection("groups")

	var groupModel models.Group
	groupModel, err := GetGroupByID(groupid, institutionID)

	if err != nil{
		return err
	}

	for studentid, exams := range(groupModel.StudentsList){
		exist := false
		pos := 0
		for i, examid := range(exams.([]string)){
			if examid == generatedExamID{
				exist = true
				pos = i
			}
		}
		if exist{
			exams.([]string)[pos] = exams.([]string)[len(exams.([]string))-1]
			exams.([]string)[len(exams.([]string))-1] = ""
			groupModel.StudentsList[studentid] = exams.([]string)[:len(exams.([]string))-1]
		}
	}
	
	updateString := bson.M{
		"$set": bson.M{
			"studentsList" : groupModel.StudentsList,
		},
	}

	id, _ := primitive.ObjectIDFromHex(groupid)

	filter := bson.M{"_id": bson.M{"$eq": id}}

	_, err = col.UpdateOne(ctx, filter, updateString)

	return err
}
