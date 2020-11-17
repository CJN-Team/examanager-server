package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Department es una estructura basica para manejar la informacion de los departamentos
type Department struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name        string   `bson:"name,omitempty" json:"name,omitempty"`
	Institution string   `bson:"institution,omitempty" json:"institution,omitempty"`
	Teachers     []string `bson:"teachers,omitempty" json:"teachers,omitempty"`
}
