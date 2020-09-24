package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Institution es una estructura basica para manejar la informacion de la institucion
type Institution struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Type      string             `bson:"type,omitempty" json:"type,omitempty"`
	Address   string             `bson:"address,omitempty" json:"address,omitempty"`
	Phone     string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Users     string             `bson:"users,omitempty" json:"users,omitempty"`
	Subjetcs  primitive.A        `bson:"subjects,omitempty" json:"subjects,omitempty"`
	Questions string             `bson:"questions,omitempty" json:"questions,omitempty"`
}
