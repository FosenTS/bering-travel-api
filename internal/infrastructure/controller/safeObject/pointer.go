package safeObject

type Pointer struct {
	Name        string  `json:"name" db:"name" binding:"required"`
	Description string  `json:"description" db:"description" binding:"required"`
	Rating      uint    `json:"rating" db:"rating" binding:"required"`
	Latitude    float64 `json:"latitude" db:"latitude" binding:"required"`
	Longitude   float64 `json:"longitude" db:"longitude" binding:"required"`
}
