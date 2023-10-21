package entity

import "github.com/achillescres/pkg/object/oid"

type Profile struct {
	UserId     oid.ID `json:"userId"`
	VisitCount uint   `json:"visitCount"`
	UserRating uint   `json:"userRating"`
	MerchCoins uint   `json:"merchCoins"`
}
