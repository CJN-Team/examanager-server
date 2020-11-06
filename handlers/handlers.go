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
	router.HandleFunc("/user", middleware.DatabaseVerify(routers.CreateUser)).Methods("POST")
	router.HandleFunc("/user", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/user", middleware.DatabaseVerify(middleware.ValidationJWT(routers.ReadUser))).Methods("GET")
	router.HandleFunc("/user", middleware.DatabaseVerify(middleware.ValidationJWT(routers.DeleteUserRouter))).Methods("DELETE")

	//Rutas para usuarios
	router.HandleFunc("/users", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetAllUsersRouter))).Methods("GET")
	router.HandleFunc("/users", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateUsersAutomatic))).Methods("POST")

	//Rutas para imagen de usuarios
	router.HandleFunc("/photo", middleware.DatabaseVerify(middleware.ValidationJWT(routers.ReadUserImage))).Methods("GET")
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

	//Rutas grupos
	router.HandleFunc("/group", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateGroup))).Methods("POST")
	router.HandleFunc("/group", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateGroup))).Methods("PUT")
	router.HandleFunc("/group", middleware.DatabaseVerify(middleware.ValidationJWT(routers.ReadGroup))).Methods("GET")
	router.HandleFunc("/group", middleware.DatabaseVerify(middleware.ValidationJWT(routers.DeleteGroup))).Methods("DELETE")

	//Rutas Grupos
	router.HandleFunc("/groups", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetAllGroups))).Methods("GET")

	//Rutas Examen
	router.HandleFunc("/exam", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateExam))).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
