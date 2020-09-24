package usersqueries

import (
	"context"
	"log"
	"time"

	dbConnection "github.com/CJN-Team/examanager-server/database"
	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetUserByEmail verifica si el usuario ya se encuentra en la base de datos por medio de el correo
func GetUserByEmail(email string) (models.User, bool, string) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	condicion := bson.M{"email": email}

	var result models.User

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	ID := result.ID.Hex()

	if error != nil {
		return result, false, ID
	}

	return result, true, ID
}

//AddUser se encarga de añadir a la base de datos un nuevo usuario
func AddUser(u models.User) (string, bool, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	u.Password, _ = PasswordEncrypt(u.Password)

	result, error := coleccion.InsertOne(contex, u)

	if error != nil {
		return "", false, error
	}

	ObjectID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjectID.String(), true, nil
}

//GetUserByID se encarga de buscar en la base de datos el usuario que posee la ID asignada
func GetUserByID(ID string) (models.User, error) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	var result models.User

	ObjectID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{"_id": ObjectID}

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	result.Password = ""

	if error != nil {
		return result, error
	}

	return result, nil
}

//GetAllUsers se encarga de traer de base de datos todos los usuarios disponibles de una categoria
func GetAllUsers(category string, page int64) ([]*models.User, bool) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	var result []*models.User

	condicion := bson.M{
		"profile": category,
	}

	searchOptions := options.Find()
	searchOptions.SetLimit(20)
	searchOptions.SetSort(bson.D{{Key: "name", Value: -1}})
	searchOptions.SetSkip((page - 1) * 20)

	pointer, error := coleccion.Find(contex, condicion, searchOptions)

	if error != nil {
		log.Fatal(error.Error())
		return result, false
	}

	for pointer.Next(context.TODO()) {
		var register models.User
		error := pointer.Decode(&register)

		if error != nil {
			return result, false
		}
		result = append(result, &register)
	}

	return result, true
}

//UpdateUser se encarga de actualizar el usuario registrado
func UpdateUser(user models.User, ID string) (bool, error) {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	userRegisterd := make(map[string]interface{})

	if len(user.Profile) > 0 {
		userRegisterd["profile"] = user.Profile
	}

	if len(user.IDType) > 0 {
		userRegisterd["idType"] = user.IDType
	}

	if len(user.Name) > 0 {
		userRegisterd["name"] = user.Name
	}

	if len(user.LastName) > 0 {
		userRegisterd["lastName"] = user.LastName
	}

	if len(user.Email) > 0 {
		userRegisterd["email"] = user.Email
	}

	userRegisterd["birthDate"] = user.BirthDate

	if len(user.Photo) > 0 {
		userRegisterd["photo"] = user.Photo
	}

	if len(user.Password) > 0 {
		userRegisterd["password"] = user.Password
	}

	updateString := bson.M{
		"$set": userRegisterd,
	}

	userID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": userID}}

	_, error := coleccion.UpdateOne(contex, filter, updateString)

	if error != nil {
		return false, error
	}

	return true, nil
}

//DeleteUser se encarga de borrar el usuario seleccionado
func DeleteUser(ID string) error {
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := dbConnection.MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	userID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": userID,
	}

	_, error := coleccion.DeleteOne(contex, condicion)

	return error
}
