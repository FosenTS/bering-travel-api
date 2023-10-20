package entity

type Pointer struct {
	Name       string     `json:"name" db:"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  " binding:"required"`
	Address    string     `json:"address" binding:"required"`
	Rating     uint       `json:"rating" binding:"required"`
	Coordinate Coordinate `json:"coordinate" binding:"required"`
}
