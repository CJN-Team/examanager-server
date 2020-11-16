package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GenerateExam es la estructura paralos examenes de cada estudiante
type GenerateExam struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Teacher     string             `bson:"teacher,omitempty" json:"teacher,omitempty"`
	Student     string             `bson:"student,omitempty" json:"student,omitempty"`
	Date        time.Time          `bson:"date,omitempty" json:"date,omitempty"`
	Topic       string             `bson:"topic,omitempty" json:"topic,omitempty"`
	Institution string             `bson:"institution,omitempty" json:"institution,omitempty"`
	Photo       string             `bson:"photo,omitempty" json:"photo,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Grade       float32            `bson:"grade,omitempty" json:"grade,omitempty"`
	Commentary  string             `bson:"commentary,omitempty" json:"commentary,omitempty"`
	Questions   map[string]float32 `bson:"question,omitempty" json:"question,omitempty"`
}
