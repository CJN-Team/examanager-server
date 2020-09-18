package models

//AnswerLogin es la estructura de respuesta de el login
type AnswerLogin struct{
	Token string `json:"token,omitempty"`
}