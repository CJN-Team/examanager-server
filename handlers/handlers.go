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

	router.HandleFunc("/user", middleware.DatabaseVerify(routers.CreateUser)).Methods("POST")
	router.HandleFunc("/user", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/user", middleware.DatabaseVerify(middleware.ValidationJWT(routers.ReadUser))).Methods("GET")
	router.HandleFunc("/login", middleware.DatabaseVerify(routers.Login)).Methods("POST")
	router.HandleFunc("/institution", middleware.DatabaseVerify(routers.InstitutionRegistration)).Methods("POST")
	router.HandleFunc("/subject", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateSubject))).Methods("POST")
	router.HandleFunc("/subject", middleware.DatabaseVerify(middleware.ValidationJWT(routers.DeleteSubject))).Methods("DELETE")
	router.HandleFunc("/subject", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateSubject))).Methods("PUT")
	router.HandleFunc("/subject", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetSubjects))).Methods("GET")
	router.HandleFunc("/users", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetAllUsersRouter))).Methods("GET")
	router.HandleFunc("/users", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateUsersAutomatic))).Methods("POST")
	router.HandleFunc("/user", middleware.DatabaseVerify(middleware.ValidationJWT(routers.DeleteUserRouter))).Methods("DELETE")
	router.HandleFunc("/questions", middleware.DatabaseVerify(middleware.ValidationJWT(routers.CreateQuestion))).Methods("POST")
	router.HandleFunc("/questions", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateQuestion))).Methods("PUT")
	router.HandleFunc("/questions", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetAllQuestionsRouter))).Methods("GET")
	router.HandleFunc("/questions", middleware.DatabaseVerify(middleware.ValidationJWT(routers.DeleteQuestionsRouter))).Methods("DELETE")
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
