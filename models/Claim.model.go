package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Claim es la estructura para proccesar los JWT recibidos
type Claim struct {
	Email string             `bson:"email" json:"email,omitempty"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Profile string			 `bson:"profile" json:"profile,omitempty"`
	Institution string		 `bson:"institution" json:"institution,omitempty"`	
	jwt.StandardClaims
}
