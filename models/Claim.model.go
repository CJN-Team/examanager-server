package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

//Claim es la estructura para proccesar los JWT recibidos
type Claim struct {
	Email       string `bson:"email" json:"email,omitempty"`
	ID          string `bson:"_id" json:"id,omitempty"`
	Profile     string `bson:"profile" json:"profile,omitempty"`
	Institution string `bson:"institution" json:"institution,omitempty"`
	jwt.StandardClaims
}
