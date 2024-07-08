package models

// Фильм
type Movie struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	RealizeDate string   `json:"realize_date"`
	Rating      int      `json:"rating"`
	Actors      []string `json:"actors" gorm:"-"`
}
