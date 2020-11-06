package models

import (
	"time"
)

//User es una estructura basica para manejar la informacion del usuario
type User struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Profile     string    `bson:"profile,omitempty" json:"profile,omitempty"`
	IDType      string    `bson:"idType,omitempty" json:"idType,omitempty"`
	Name        string    `bson:"name,omitempty" json:"name,omitempty"`
	LastName    string    `bson:"lastName,omitempty" json:"lastName,omitempty"`
	Email       string    `bson:"email,omitempty" json:"email"`
	BirthDate   time.Time `bson:"birthDate,omitempty" json:"birthDate,omitempty"`
	Photo       string    `bson:"photo,omitempty" json:"photo,omitempty"`
	Institution string    `bson:"institution" json:"institution,omitempty"`
	Password    string    `bson:"password,omitempty" json:"password,omitempty"`
	
}
