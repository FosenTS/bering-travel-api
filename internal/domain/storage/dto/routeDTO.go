package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
)

type RouteCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	AirlCode             string   `json:"airCode" db:"air_code" binding:"required"`
	Number               string   `json:"number" db:"number" binding:"required"`
	OriginIATA           string   `json:"originIATA" db:"origin_iata" binding:"required"`
	DestinationIATA      string   `json:"destinationIATA" db:"destination_iata" binding:"required"`
	Distance             *float64 `json:"distance" db:"distance" binding:"required"`
	TemporaryConsumption *float64 `json:"temporaryConsumption" db:"temporary_consumption" binding:"required"`
	ConstantConsumption  *float64 `json:"constantConsumption" db:"constant_consumption" binding:"required"`
}

func NewRouteCreate(
	airlCode,
	number,
	originIATA,
	destinationIATA string,
	distance *float64,
	temporaryConsumption *float64,
	constantConsumption *float64,
) *RouteCreate {
	return &RouteCreate{
		AirlCode:             airlCode,
		Number:               number,
		OriginIATA:           originIATA,
		DestinationIATA:      destinationIATA,
		Distance:             distance,
		TemporaryConsumption: temporaryConsumption,
		ConstantConsumption:  constantConsumption,
	}
}

func (rC *RouteCreate) ToEntity(id oid.ID) *entity.Route {
	return entity.NewRoute(
		id,
		rC.AirlCode,
		rC.Number,
		rC.OriginIATA,
		rC.DestinationIATA,
		rC.Distance,
		rC.TemporaryConsumption,
		rC.ConstantConsumption,
	)
}
