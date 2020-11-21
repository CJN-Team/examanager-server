package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GenerateExam es la estructura paralos examenes de cada estudiante
type GenerateExam struct {
	ID              primitive.ObjectID       `bson:"_id,omitempty" json:"id"`
	View            bool                     `bson:"view,omitempty" json:"view,omitempty"`
	State           bool                     `bson:"state,omitempty" json:"state,omitempty"`
	MockExam        bool                     `bson:"mockExam,omitempty" json:"mockExam,omitempty"`
	Finish          bool                     `bson:"finish,omitempty" json:"finish,omitempty"`
	TeacherID       string                   `bson:"teacherid,omitempty" json:"teacherid,omitempty"`
	Teacher         string                   `bson:"teacher,omitempty" json:"teacher,omitempty"`
	Student         string                   `bson:"student,omitempty" json:"student,omitempty"`
	StudentID       string                   `bson:"studentid,omitempty" json:"studentid,omitempty"`
	Date            time.Time                `bson:"date,omitempty" json:"date,omitempty"`
	Topic           string                   `bson:"topic,omitempty" json:"topic,omitempty"`
	InstitutionID   string                   `bson:"institutionid,omitempty" json:"institutionid,omitempty"`
	InstitutionName string                   `bson:"institutionname,omitempty" json:"institutionname,omitempty"`
	Photo           string                   `bson:"photo,omitempty" json:"photo,omitempty"`
	Name            string                   `bson:"name,omitempty" json:"name,omitempty"`
	Grade           float64                  `bson:"grade" json:"grade"`
	Commentary      string                   `bson:"commentary,omitempty" json:"commentary,omitempty"`
	Questions       map[string][]interface{} `bson:"question,omitempty" json:"question,omitempty"`
}
