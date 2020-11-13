package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoConexion es el objeto para conectarme a la base de datos
var MongoConexion = ConnectBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://" + os.Getenv("USER") + ":" + os.Getenv("PASSWORD") + "@examangerdb.1dbzg.mongodb.net/" + os.Getenv("DATABASE") + "?retryWrites=true&w=majority")

//ConnectBD es la funcion que me permite conectarme con la base de datos
func ConnectBD() *mongo.Client {
	client, error := mongo.Connect(context.TODO(), clientOptions)

	if error != nil {
		log.Fatal(error.Error())

		return client
	}

	error = client.Ping(context.TODO(), nil)

	if error != nil {
		log.Fatal(error.Error())

		return client
	}

	log.Println("Conexion Exitosa MongoDB")
	return client
}

//CheckConnection permite comprobar si la base de datos esta activa
func CheckConnection() int {
	error := MongoConexion.Ping(context.TODO(), nil)

	if error != nil {
		return 0
	}
	return 1
}
