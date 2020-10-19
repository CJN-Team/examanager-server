package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Group es una estructura basica para manejar la informacion de los grupos
type Group struct {
	ID           string      `bson:"_id,omitempty" json:"id"`
	Name         string      `bson:"name,omitempty" json:"name,omitempty"`
	StudentsList primitive.M `bson:"studentsList,omitempty" json:"studentsList,omitempty"`
	Teacher      string      `bson:"teacher,omitempty" json:"teacher,omitempty"`
}
