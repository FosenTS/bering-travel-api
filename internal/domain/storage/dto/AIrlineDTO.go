package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
)

type AirlineCreate struct {
	Create `json:"-" db:"-" binding:"-"`
	Name   string `json:"name" db:"name" binding:"required"`
	Code   string `json:"code" db:"code" binding:"required"`
}

func NewAirlineCreate(name string, code string) *AirlineCreate {
	return &AirlineCreate{Name: name, Code: code}
}

func (aC *AirlineCreate) ToEntity(id oid.ID) *entity.Airline {
	return entity.NewAirline(id, aC.Name, aC.Code)
}
