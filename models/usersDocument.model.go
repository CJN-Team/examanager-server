package models

//SearchDocument se encarga de almacenar el link del documento a buscar
type SearchDocument struct {
	DocumentLink string `json:"documentLink,omitempty"`
	UserProfile  string `json:"userProfile,omitempty"`
}
