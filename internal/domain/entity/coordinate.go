package entity

type Coordinate struct {
	Latitude  float64 `json:"latitude" db:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" db:"longitude" binding:"required"`
}
