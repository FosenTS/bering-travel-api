package entity

type User_visit struct {
	Id        uint   `json:"id" db:"id" binding:"required"`
	UserId    uint   `json:"userId" db:"user_id" binding:"required"`
	PointerId uint   `json:"pointerId" db:"pointer_id" binding:"required"`
	Rating    uint   `json:"rating" db:"rating" binding:"required"`
	Comment   string `json:"comment" db:"comment" binding:"required"`
}
