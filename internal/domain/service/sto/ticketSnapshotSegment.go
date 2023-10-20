package sto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"time"
)

type TicketSnapshotsSegment struct {
	FlightID oid.ID                   `json:"flightID" db:"flight_id" binding:"required"`
	Time     time.Time                `json:"time" db:"time" binding:"required"`
	Tickets  []*entity.TicketSnapshot `json:"tickets" db:"tickets" binding:"required"`
}

func NewTicketSnapshotSegment(flightID oid.ID, time time.Time, tickets []*entity.TicketSnapshot) *TicketSnapshotsSegment {
	return &TicketSnapshotsSegment{FlightID: flightID, Time: time, Tickets: tickets}
}
