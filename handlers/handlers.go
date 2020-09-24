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
	router.HandleFunc("/createInstitution", middleware.DatabaseVerify(routers.InstitutionRegistration)).Methods("POST")
	router.HandleFunc("/login", middleware.DatabaseVerify(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middleware.DatabaseVerify(middleware.ValidationJWT(routers.ReadUser))).Methods("GET")
	router.HandleFunc("/UpdateUser", middleware.DatabaseVerify(middleware.ValidationJWT(routers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/Users", middleware.DatabaseVerify(middleware.ValidationJWT(routers.GetAllUsersRouter))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
