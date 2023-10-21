package entity

import "time"

type Commit struct {
	UserID uint      `json:"userId" db:"user_id" binding:"required"`
	Commit string    `json:"comment" db:"comment" binding:"required"`
	Time   time.Time `json:"time" db:"time" binding:"required"`
}
