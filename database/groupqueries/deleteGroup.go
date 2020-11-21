package groupqueries

import (
	"fmt"
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

//DeleteExamsOfStudents elimina los examanes generados del padre de los estudiantes de un grupo, metodo de Andres
func DeleteExamsOfStudents(groupid string, generatedExamID string, loggedUser string,institutionID string)(error) {

	newStudentsList := bson.M{}

	var groupModel models.Group
	groupModel, err := GetGroupByID(groupid, institutionID)

	if err != nil{
		return err
	}

	for studentid, exams := range(groupModel.StudentsList){

		newExams := bson.A{}
		j:=0
		for _, examid := range(exams.(primitive.A)){

			if examid.(string) != generatedExamID{
				newExams = append(newExams,examid.(string))
				j++
			}

		}

		newStudentsList[studentid] = newExams
	}

	groupModel.StudentsList = newStudentsList
	_,err = UpdateGroup(groupModel,groupModel.ID,loggedUser,institutionID)
	fmt.Println(err)
	return err
}
