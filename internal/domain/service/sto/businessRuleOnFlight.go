package sto

import "github.com/achillescres/pkg/object/oid"

type BusinessRuleOnFlight struct {
	ID        oid.ID `json:"id" db:"id" binding:"required"`
	Name      string `json:"name" db:"name" binding:"required"`
	Rate      string `json:"rate" binding:"required"`
	GroupID   oid.ID `json:"groupID" db:"id" binding:"required"`
	GroupName string `json:"groupName" db:"name" binding:"required"`
}

func NewBusinessRuleOnFlight(ID oid.ID, name string, rate string, GroupID oid.ID, GroupName string) *BusinessRuleOnFlight {
	return &BusinessRuleOnFlight{ID: ID, Name: name, Rate: rate, GroupID: GroupID, GroupName: GroupName}
}
