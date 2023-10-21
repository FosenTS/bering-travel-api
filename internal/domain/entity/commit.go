package entity

type Commit struct {
	UserID uint   `json:"userId" db:"user_id" binding:"required"`
	Commit string `json:"comment" db:"comment" binding:"required"`
}

func NewCommit(userID uint, commit string) *Commit {
	return &Commit{UserID: userID, Commit: commit}
}
