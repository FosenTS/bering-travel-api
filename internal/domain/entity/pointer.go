package entity

import "time"

type Pointer struct {
	Id          uint      `json:"id" db:"id" binding:"required"`
	Name        string    `json:"name" db:"name" binding:"required"`
	Description string    `json:"description" db:"description" binding:"required"`
	Rating      uint      `json:"rating" db:"rating" binding:"required"`
	Latitude    float64   `json:"latitude" db:"latitude" binding:"required"`
	Longitude   float64   `json:"longitude" db:"longitude" binding:"required"`
	Time        time.Time `json:"time" db:"time" binding:"required"`
}
