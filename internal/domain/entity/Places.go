package entity

type Places struct {
	Id          uint      `json:"id" db:"id" binding:"required"`
	Name        string    `json:"name" db:"name" binding:"required"`
	Description string    `json:"description" db:"description" binding:"required"`
	Rating      uint      `json:"rating" db:"rating" binding:"required"`
	Latitude    float64   `json:"latitude" db:"latitude" binding:"required"`
	Longitude   float64   `json:"longitude" db:"longitude" binding:"required"`
	Comments    []*Commit `json:"comments"`
}

func NewPlaces(id uint, name string, description string, rating uint, latitude float64, longitude float64, comments []*Commit) *Places {
	return &Places{Id: id, Name: name, Description: description, Rating: rating, Latitude: latitude, Longitude: longitude, Comments: comments}
}
