package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Exam es una estructura basica para el manejo de los examenes
type Exam struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	GroupID         string             `bson:"groupId,omitempty" json:"groupId"`
	Name            string             `bson:"name,omitempty" json:"name,omitempty"`
	NumberQuestions int                `bson:"numberQuestions,omitempty" json:"numberQuestions,omitempty"`
	Difficulty      int                `bson:"difficulty,omitempty" json:"difficulty,omitempty"`
	TopicQuestion   string             `bson:"topicQuestion,omitempty" json:"topicQuestion,omitempty"`
	Date            time.Time          `bson:"date,omitempty" json:"date,omitempty"`
}
