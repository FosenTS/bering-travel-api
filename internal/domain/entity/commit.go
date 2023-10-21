package entity

import "github.com/achillescres/pkg/object/oid"

type Commit struct {
	UserID oid.ID `json:"userId" binding:"required"`
	Commit string `json:"commit" binding:"required"`
}

func NewCommit(userID oid.ID, commit string) *Commit {
	return &Commit{UserID: userID, Commit: commit}
}
