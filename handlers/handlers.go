package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/CJN-Team/examanager-server/middleware"
	"github.com/CJN-Team/examanager-server/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Manejadores setea al handler y pone a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	//Rutas para usuario
	router.HandleFunc("/admin", middleware.DatabaseVerify(routers.CreateUser)).Methods("POST")

	router.HandleFunc("/user", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateUser))).Methods("POST")
	router.HandleFunc("/user", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/user", middleware.DatabaseVerify(middleware.ValidationJWT(routers.ReadUser))).Methods("GET")
	router.HandleFunc("/user", middleware.DatabaseVerify(middleware.ValidationJWT(routers.DeleteUserRouter))).Methods("DELETE")

	//Rutas para usuarios
	router.HandleFunc("/users", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetAllUsersRouter))).Methods("GET")
	router.HandleFunc("/users", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateUsersAutomatic))).Methods("POST")

	//Rutas para imagen de usuarios
	router.HandleFunc("/photo", middleware.DatabaseVerify(routers.ReadUserImage)).Methods("GET")
	router.HandleFunc("/photo", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UploadUserImage))).Methods("PUT")

	//Inicio de sesion
	router.HandleFunc("/login", middleware.DatabaseVerify(routers.Login)).Methods("POST")

	//Rutas institucion
	router.HandleFunc("/institution", middleware.DatabaseVerify(routers.InstitutionRegistration)).Methods("POST")

	//Rutas asignaturas
	router.HandleFunc("/subject", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateSubject))).Methods("POST")
	router.HandleFunc("/subject", middleware.DatabaseVerify(middleware.ValidationJWT(routers.DeleteSubject))).Methods("DELETE")
	router.HandleFunc("/subject", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateSubject))).Methods("PUT")
	router.HandleFunc("/subject", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetSubjects))).Methods("GET")

	//Rutas Preguntas
	router.HandleFunc("/questions", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateQuestion))).Methods("POST")
	router.HandleFunc("/questions", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateQuestion))).Methods("PUT")
	router.HandleFunc("/questions", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetAllQuestionsRouter))).Methods("GET")
	router.HandleFunc("/questions", middleware.DatabaseVerify(middleware.ValidationJWT(routers.DeleteQuestionsRouter))).Methods("DELETE")
	router.HandleFunc("/getonequestion", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetOnequestion))).Methods("GET")

	//Rutas grupos
	router.HandleFunc("/group", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateGroup))).Methods("POST")
	router.HandleFunc("/group", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateGroup))).Methods("PUT")
	router.HandleFunc("/group", middleware.DatabaseVerify(middleware.ValidationJWT(routers.ReadGroup))).Methods("GET")
	router.HandleFunc("/group", middleware.DatabaseVerify(middleware.ValidationJWT(routers.DeleteGroup))).Methods("DELETE")

	//Rutas Grupos
	router.HandleFunc("/groups", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetAllGroups))).Methods("GET")
	router.HandleFunc("/groupProgress", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetAllWatchedTopics))).Methods("GET")
	router.HandleFunc("/groupUserGradesAll", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetUserGradesAllGroups))).Methods("GET")
	router.HandleFunc("/groupUserGrades", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetUserGrades))).Methods("GET")

	//Rutas departamento
	router.HandleFunc("/departament", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateDepartment))).Methods("POST")
	router.HandleFunc("/departament", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateDepartment))).Methods("PUT")
	router.HandleFunc("/departament", middleware.DatabaseVerify(middleware.ValidationJWT(routers.ReadDepartment))).Methods("GET")
	router.HandleFunc("/departament", middleware.DatabaseVerify(middleware.ValidationJWT(routers.DeleteDepartment))).Methods("DELETE")

	//Rutas departamentos
	router.HandleFunc("/departaments", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetAllDepartments))).Methods("GET")

	//Rutas Examen
	router.HandleFunc("/exam", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateExam))).Methods("POST")
	router.HandleFunc("/exam", middleware.DatabaseVerify(middleware.ValidationJWT(routers.DeleteExam))).Methods("DELETE")
	router.HandleFunc("/exam", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateExam))).Methods("PUT")
	router.HandleFunc("/exam", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetAllExams))).Methods("GET")
	router.HandleFunc("/generatedexam", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateExamGrade))).Methods("PUT")
	router.HandleFunc("/generatedexam", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetGenerateExam))).Methods("GET")
	router.HandleFunc("/generatedexam", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GradeExam))).Methods("POST")
	router.HandleFunc("/examgenerator", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateGenerateExam))).Methods("PUT")
	router.HandleFunc("/exampdf", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GeneratePDF))).Methods("PUT")
	router.HandleFunc("/exampdf", middleware.DatabaseVerify(routers.DownloadPDF)).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
