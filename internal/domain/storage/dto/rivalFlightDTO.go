package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"time"
)

type RivalFlightCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	AirlCode              string    `json:"airlCode" db:"airl_code" binding:"required"`
	ObservationTime       time.Time `json:"observationTime" db:"observation_time" binding:"required"`
	Origin                string    `json:"origin" db:"origin" binding:"required"`
	Destination           string    `json:"destination" db:"destination" binding:"required"`
	IsOneWay              int       `json:"isOneWay" db:"is_one_way" binding:"required"`
	OutboundFlightNumber  string    `json:"outboundFlightNumber" db:"outbound_flight_number" binding:"required"`
	InboundFlightNumber   string    `json:"inboundFlightNumber" db:"inbound_flight_number" binding:"required"`
	OutboundDepartureTime time.Time `json:"outboundDepartureTime" db:"outbound_departure_time" binding:"required"`
	OutboundArrivalTime   time.Time `json:"outboundArrivalTime" db:"outbound_arrival_time" binding:"required"`
	InboundDepartureTime  time.Time `json:"inboundDepartureTime" db:"inbound_departure_time" binding:"required"`
	InboundArrivalTime    time.Time `json:"InboundArrivalTime" db:"inbound_arrival_time" binding:"required"`
	Price                 float32   `json:"price" db:"price" binding:"required"`
	Tax                   float32   `json:"tax" db:"tax" binding:"required"`
	Currency              string    `json:"currency" db:"currency" binding:"required"`
}

func NewRivalFlightCreate(
	airlCode string,
	observationTime time.Time,
	origin string,
	destination string,
	isOneWay int,
	outboundFlightNumber string,
	inboundFlightNumber string,
	outboundDepartureTime time.Time,
	outboundArrivalTime time.Time,
	inboundDepartureTime time.Time,
	inboundArrivalTime time.Time,
	price float32,
	tax float32,
	currency string,
) *RivalFlightCreate {
	return &RivalFlightCreate{
		AirlCode:              airlCode,
		ObservationTime:       observationTime,
		Origin:                origin,
		Destination:           destination,
		IsOneWay:              isOneWay,
		OutboundFlightNumber:  outboundFlightNumber,
		InboundFlightNumber:   inboundFlightNumber,
		OutboundDepartureTime: outboundDepartureTime,
		OutboundArrivalTime:   outboundArrivalTime,
		InboundDepartureTime:  inboundDepartureTime,
		InboundArrivalTime:    inboundArrivalTime,
		Price:                 price,
		Tax:                   tax,
		Currency:              currency,
	}
}

func (rfC *RivalFlightCreate) ToEntity(id oid.ID) *entity.RivalFlight {
	return entity.NewRivalFlight(
		id,
		rfC.AirlCode,
		rfC.ObservationTime,
		rfC.Origin,
		rfC.Destination,
		rfC.IsOneWay,
		rfC.InboundFlightNumber,
		rfC.OutboundFlightNumber,
		rfC.OutboundDepartureTime,
		rfC.OutboundArrivalTime,
		rfC.InboundDepartureTime,
		rfC.InboundArrivalTime,
		rfC.Price,
		rfC.Tax,
		rfC.Currency,
	)
}
