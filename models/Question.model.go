package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Question es una estructura basica para manejar las preguntas de la aplicacion
type Question struct {
	ID        	primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	Thematic	string				`bson:"thematic,omitempty" json:"thematic,omitempty"`
	Subject		string				`bson:"subject,omitempty" json:"subject,omitempty"`
	Pregunta	string				`bson:"pregunta,omitempty" json:"pregunta,omitempty"`
	Category	string				`bson:"category,omitempty" json:"category,omitempty"`
	Options   	[]string			`bson:"options,omitempty" json:"options,omitempty"`
	Answer		string				`bson:"Answer,omitempty" json:"Answer,omitempty"`
	Difficulty	int					`bson:"difficulty,omitempty" json:"difficulty,omitempty"`
}
