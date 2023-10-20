package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"time"
)

type FLightCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	RouteID          oid.ID `json:"routeID" db:"route_id" binding:"required"`
	AirlCode         string `json:"airlCode" db:"airl_code" binding:"required"`
	Number           string `json:"number" db:"number" binding:"required"`
	OrigIATA         string `json:"origIata" db:"orig_iata" binding:"required"`
	DestIATA         string `json:"destIata" db:"dest_iata" binding:"required"`
	DepartureTime    string `json:"departureTime" db:"departure_time" binding:"required"`
	ArrivalTime      string `json:"arrivalTime" db:"arrival_time" binding:"required"`
	OriginOrder      int    `json:"originOrder" db:"orig_order" binding:"required"`
	DestinationOrder int    `json:"destinationOrder" db:"dest_order" binding:"required"`

	FltDate   string  `json:"fltDate" db:"flt_date" binding:"required"`
	Aircraft  string  `json:"aircraft" db:"aircraft" binding:"required"`
	Resource  int     `json:"resource" db:"resource" binding:"required"`
	TotalCash float64 `json:"totalCash" db:"total_cash" binding:"required"`

	IsComposite bool `json:"isComposite" db:"is_composite" binding:"required"`

	State           string    `json:"state" db:"state" binding:"required"`
	LastUpdate      time.Time `json:"lastUpdate" db:"last_update" binding:"required"`
	CorrectlyParsed bool      `json:"correctlyParsed" db:"correctly_parsed" binding:"required"`
}

func NewFLightCreate(
	routeID oid.ID,
	airlCode string,
	number string,
	origIATA string,
	destIATA string,
	departureTime string,
	arrivalTime string,
	originOrder int,
	destinationOrder int,
	fltDate string,
	aircraft string,
	resource int,
	totalCash float64,
	isComposite bool,
	state string,
	lastUpdate time.Time,
	correctlyParsed bool,
) *FLightCreate {
	return &FLightCreate{
		RouteID:          routeID,
		AirlCode:         airlCode,
		Number:           number,
		OrigIATA:         origIATA,
		DestIATA:         destIATA,
		DepartureTime:    departureTime,
		ArrivalTime:      arrivalTime,
		OriginOrder:      originOrder,
		DestinationOrder: destinationOrder,
		FltDate:          fltDate,
		Aircraft:         aircraft,
		Resource:         resource,
		TotalCash:        totalCash,
		IsComposite:      isComposite,
		State:            state,
		LastUpdate:       lastUpdate,
		CorrectlyParsed:  correctlyParsed,
	}
}

func (fC *FLightCreate) ToEntityWithFilledAdditionalFields(id oid.ID) *entity.Flight {
	return entity.NewFlight(
		id,
		fC.RouteID,
		fC.AirlCode,
		fC.Number,
		fC.OrigIATA,
		fC.DestIATA,
		fC.DepartureTime,
		fC.ArrivalTime,
		fC.OriginOrder,
		fC.DestinationOrder,
		fC.FltDate,
		fC.Aircraft,
		fC.Resource,
		fC.TotalCash,
		0,
		0,
		fC.IsComposite,
		fC.State,
		fC.LastUpdate,
		fC.CorrectlyParsed,
	)
}
