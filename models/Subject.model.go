package models

//Subject es la estructura para ver una de las materias registadas en una institucion
type Subject struct {
	Name   string   `bson:"name,omitempty" json:"name,omitempty"`
	Topics []string `bson:"topics,omitempty" json:"topics,omitempty"`
}
