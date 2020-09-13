package main

import (
	"log"

	"github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/handlers"
)

func main() {
	if database.CheckConnection() == 0 {
		log.Fatal("Sin conexion a mongoDB")
		return
	}
	handlers.Manejadores()
}
