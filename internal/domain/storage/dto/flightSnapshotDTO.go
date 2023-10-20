package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"time"
)

type FlightSnapshotCreate struct {
	Create   `json:"-" db:"-" binding:"-"`
	FlightID oid.ID    `json:"flightID" db:"flight_id" binding:"required"`
	Time     time.Time `json:"time" db:"time" binding:"required"`

	TotalCash float64 `json:"totalCash" db:"total_cash" binding:"required"`
	Load      int     `json:"load" db:"load" binding:"required"`
	ARR       float64 `json:"ARR" db:"ARR" binding:"required"`
}

func NewFlightSnapshotCreate(flightID oid.ID, time time.Time, totalCash float64, load int, ARR float64) *FlightSnapshotCreate {
	return &FlightSnapshotCreate{
		FlightID:  flightID,
		Time:      time,
		TotalCash: totalCash,
		Load:      load,
		ARR:       ARR,
	}
}

func (fSC *FlightSnapshotCreate) ToEntity(id oid.ID) *entity.FlightSnapshot {
	return entity.NewFlightSnapshot(
		id,
		fSC.FlightID,
		fSC.Time,
		fSC.TotalCash,
		fSC.Load,
		fSC.ARR,
	)
}

func NewFlightSnapshotCreateFromFlight(f entity.Flight, t time.Time) *FlightSnapshotCreate {
	return NewFlightSnapshotCreate(
		f.ID,
		t,
		f.TotalCash,
		f.Load,
		f.ARR,
	)
}
