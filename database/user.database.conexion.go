package database

import (
	"context"
	"time"


	"github.com/CJN-Team/examanager-server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetUserByEmail verifica si el usuario ya se encuentra en la base de datos por medio de el correo
func GetUserByEmail(email string) (models.User, bool, string) {

	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("examanager_db")

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

	database := MongoConexion.Database("examanager_db")

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

	database := MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	var result models.User

	ObjectID,_ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{"_id": ObjectID}

	error := coleccion.FindOne(contex, condicion).Decode(&result)

	result.Password=""

	if error != nil {
		return result, error
	}

	return result, nil
}

//UpdateUser se encarga de actualizar el usuario registrado
func UpdateUser(user models.User, ID string) (bool, error){
	contex, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConexion.Database("examanager_db")

	coleccion := database.Collection("users")

	userRegisterd := make(map[string]interface{})

	if len(user.Profile) > 0{
		userRegisterd["profile"]=user.Profile
	}

	if len(user.IDType) > 0{
		userRegisterd["idType"]=user.IDType
	}

	if len(user.Name) > 0{
		userRegisterd["name"]=user.Name
	}

	if len(user.LastName) > 0{
		userRegisterd["lastName"]=user.LastName
	}

	if len(user.Email) > 0{
		userRegisterd["email"]=user.Email
	}

	userRegisterd["birthDate"]=user.BirthDate

	if len(user.Photo) > 0{
		userRegisterd["photo"]=user.Photo
	}

	if len(user.Password) > 0{
		userRegisterd["password"]=user.Password
	}

	updateString:= bson.M{
		"$set":userRegisterd,
	}

	userID,_:= primitive.ObjectIDFromHex(ID)

	filter:= bson.M{"_id":bson.M{"$eq":userID}}

	_,error := coleccion.UpdateOne(contex,filter,updateString)

	if error!= nil{
		return false,error
	}

	return true , nil
}