package models


//Subject es una estructura basica para las asignaturas creadas por las instituciones
type Subject struct {
	Name       string             `bson:"name,omitempty" json:"name,omitempty"`
	TopicsList []string           `bson:"topicsList,omitempty" json:"topicsList,omitempty"`
}
