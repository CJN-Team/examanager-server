package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//UsersXInstitution es una estructura basica para manejar los usuarios registrados en una institucion
type UsersXInstitution struct {
	ID        			primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	InstitutionName     string             `bson:"institutionName,omitempty" json:"institutionName,omitempty"`
	AdminsList     		primitive.A        `bson:"adminsList" json:"adminsList"` //primitive.A stands for BSON Array in MongoDB
	StudentsList     	primitive.A        `bson:"studentsList" json:"studentsList"`
	TeachersList     	primitive.A        `bson:"teachersList" json:"teachersList"`
}
