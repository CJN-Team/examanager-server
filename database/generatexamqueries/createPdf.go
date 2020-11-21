package generatexamqueries

import (
	"context"
	"log"
	"strconv"
	"time"

	//dbExam"github.com/CJN-Team/examanager-server/database/examqueries"
	dbConnection "github.com/CJN-Team/examanager-server/database"
	questionsDB "github.com/CJN-Team/examanager-server/database/questionsqueries"
	"github.com/CJN-Team/examanager-server/models"
	"github.com/signintech/gopdf"
	"go.mongodb.org/mongo-driver/bson"
)

// CreatePDF creara los pdf de los examenes
func CreatePDF(exams models.Exam, institution string) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	err := pdf.AddTTFFont("DROID", "DroidSerif-Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("DROID", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	for i := 0; i < len(exams.GenerateExam); i++ {
		pdf.AddPage()
		x, y := 30.0, 40.0
		generate, _ := GetGenerateExamByID(exams.GenerateExam[i], institution)
		user, _ := getUserByIDOneInstitution(generate.StudentID, institution)
		pdf.Image("assets/users/"+user.Photo, 425, 30, nil)
		pdf.SetX(x)
		pdf.SetY(y)
		pdf.Text(generate.InstitutionName)
		y = y + 20
		pdf.SetX(x)
		pdf.SetY(y)
		pdf.Text("Alumno: " + generate.Student)
		y = y + 20
		pdf.SetX(x)
		pdf.SetY(y)
		pdf.Text(generate.Name)
		y = y + 20
		pdf.SetX(x)
		pdf.SetY(y)
		pdf.Text(exams.SubjectID)
		y = y + 20
		pdf.SetX(x)
		pdf.SetY(y)
		pdf.Text("Profesor: " + generate.Teacher)
		y = y + 50
		num := 1
		page := 0
		for key := range generate.Questions {
			if page == 2 {
				y = 40
				page = 0
				pdf.AddPage()
				question, _, _ := questionsDB.GetQuestionByID(key, institution)

				pdf.SetX(x)
				pdf.SetY(y)
				s := strconv.Itoa(num)
				pdf.Text(s + ". " + question.Pregunta + " (" + question.Category + ")")
				num++
				for j := 0; j < len(question.Options); j++ {
					if question.Options[j] == "Abierta" {
						break
					}
					y = y + 20
					pdf.SetX(x + 10)
					pdf.SetY(y)
					pdf.Text(string(j+97) + ". " + question.Options[j])

				}
				y = 400
			} else {
				page = page + 1
				question, _, _ := questionsDB.GetQuestionByID(key, institution)
				pdf.SetX(x)
				pdf.SetY(y)
				s := strconv.Itoa(num)
				pdf.Text(s + ". " + question.Pregunta + " (" + question.Category + ")")
				num++
				for j := 0; j < len(question.Options); j++ {
					if question.Options[j] == "Abierta" {
						break
					}
					y = y + 20
					pdf.SetX(x + 10)
					pdf.SetY(y)
					pdf.Text(string(j+97) + ". " + question.Options[j])

				}
				y = y + 240
			}
		}
	}
	name := exams.ID.Hex()

	pdf.WritePdf("exam-pdf/" + name + ".pdf")
}

//GetUserByIDOneInstitution se encarga de buscar en la base de datos el usuario que posee la ID asignada en una sola institucion
func getUserByIDOneInstitution(ID string, institution string) (models.User, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	var result models.User

	ObjectID := ID

	condicion := bson.M{"_id": ObjectID, "institution": institution}

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	result.Password = ""

	if error != nil {
		return result, error
	}

	return result, nil
}
