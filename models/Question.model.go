package models

//Question es una estructura basica para manejar las preguntas de la aplicacion
type Question struct {
	ID         string			`bson:"_id,omitempty" json:"id"`
	Topic      string			`bson:"topic,omitempty" json:"topic,omitempty"`
	Subject    string			`bson:"subject,omitempty" json:"subject,omitempty"`
	Pregunta   string			`bson:"question,omitempty" json:"question,omitempty"`
	Category   string			`bson:"category,omitempty" json:"category,omitempty"`
	Options    []string			`bson:"options,omitempty" json:"options,omitempty"`
	Answer     []int			`bson:"answer,omitempty" json:"answer,omitempty"`
	Difficulty int				`bson:"difficulty,omitempty" json:"difficulty,omitempty"`
}
