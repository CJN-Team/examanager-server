package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Exam es una estructura basica para el manejo de los examenes
type Exam struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	View          bool               `bson:"view,omitempty" json:"view,omitempty"`
	MockExam      bool               `bson:"mockExam,omitempty" json:"mockExam,omitempty"`
	Institution   string             `bson:"institution,omitempty" json:"institution,omitempty"`
	SubjectID     string             `bson:"subjectID,omitempty" json:"subjectID"`
	GroupID       string             `bson:"groupId,omitempty" json:"groupId"`
	Name          string             `bson:"name,omitempty" json:"name,omitempty"`
	State         bool               `bson:"state,omitempty" json:"state,omitempty"`
	Difficulty    []int              `bson:"difficulty,omitempty" json:"difficulty,omitempty"` //posicion 0 dificultad 1, posicion 1 dificultad 2, pocicion 2 dificultad 3
	TopicQuestion string             `bson:"topicQuestion,omitempty" json:"topicQuestion,omitempty"`
	Date          time.Time          `bson:"date,omitempty" json:"date,omitempty"`
	GenerateExam  []string           `bson:"generateExam,omitempty" json:"generateExam,omitempty"`
}
