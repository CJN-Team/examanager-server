package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GenerateExam es la estructura paralos examenes de cada estudiante
type GenerateExam struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Teacher         string             `bson:"teacher,omitempty" json:"teacher,omitempty"`
	Student         string             `bson:"student,omitempty" json:"student,omitempty"`
	StudentID       string             `bson:"studentid,omitempty" json:"studentid,omitempty"`
	Date            time.Time          `bson:"date,omitempty" json:"date,omitempty"`
	Topic           string             `bson:"topic,omitempty" json:"topic,omitempty"`
	InstitutionID   string             `bson:"institutionid,omitempty" json:"institutionid,omitempty"`
	InstitutionName string             `bson:"institutionname,omitempty" json:"institutionname,omitempty"`
	Photo           string             `bson:"photo,omitempty" json:"photo,omitempty"`
	Name            string             `bson:"name,omitempty" json:"name,omitempty"`
	Grade           float32            `bson:"grade" json:"grade"`
	Commentary      string             `bson:"commentary,omitempty" json:"commentary,omitempty"`
	Questions       map[string]float32 `bson:"question,omitempty" json:"question,omitempty"`
}
