package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//QuestionsXInstitution es una estructura basica para manejar las preguntas de una institucion
type QuestionsXInstitution struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	InstitutionName string             `bson:"institutionName,omitempty" json:"institutionName,omitempty"`
	QuestionsList 	primitive.A 	   `bson:"questionsList" json:"questionsList"`
	//primitive.A stands for BSON Array in MongoDB
}
