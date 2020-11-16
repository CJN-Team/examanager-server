package models

//Department es una estructura basica para manejar la informacion de los departamentos
type Department struct {
	ID          string   `bson:"_id,omitempty" json:"id"`
	Name        string   `bson:"name,omitempty" json:"name,omitempty"`
	Institution string   `bson:"institution,omitempty" json:"institution,omitempty"`
	Teachers     []string `bson:"teachers,omitempty" json:"teachers,omitempty"`
}
