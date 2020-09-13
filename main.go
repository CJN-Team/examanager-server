package main

import (
	"server/database"
	"server/handlers"
	"log"
)

func main() {
	if database.CheckConnection() == 0 {
		log.Fatal("Sin conexion a mongoDB")
		return
	}
	handlers.Manejadores()
}
