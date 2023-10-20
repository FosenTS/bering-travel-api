package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"time"
)

type TicketSnapshotCreate struct {
	Create `json:"-" db:"_" binding:"required"`

	Time             time.Time `json:"time" db:"time" binding:"required"`
	FlightID         oid.ID    `json:"flightID" db:"flight_id" binding:"required"`
	FlightSnapshotID oid.ID    `json:"flightSnapshotID" db:"flight_snapshot_id" binding:"required"`

	Cabin           string    `json:"cabin" db:"cabin" binding:"required"`
	Resource        int       `json:"resource" db:"resource" binding:"required"`
	Class           string    `json:"class" db:"class" binding:"required"`
	BookedSeats     int       `json:"bookedSeats" db:"booked_seats" binding:"required"`
	NestingCapacity int       `json:"nestingCapacity" db:"nesting_capacity" binding:"required"`
	AvailableSeats  int       `json:"availableSeats" db:"available_seats" binding:"required"`
	TotalCash       float64   `json:"totalCash" db:"total_cash" binding:"required"`
	LastBooking     time.Time `json:"lastBooking" db:"last_booking" binding:"required"`

	CorrectlyParsed bool `json:"correct" db:"correctly_parsed" binding:"required"`

	Noshow *int `json:"noshow" db:"noshow" binding:"required"`
}

func NewTicketSnapshotCreate(
	time time.Time,
	flightID oid.ID,
	flightSnapshotID oid.ID,
	cabin string,
	resource int,
	class string,
	bookedSeats int,
	nestingCapacity int,
	availableSeats int,
	totalCash float64,
	lastBooking time.Time,
	correctlyParsed bool,
	noshow *int,
) *TicketSnapshotCreate {
	return &TicketSnapshotCreate{
		Time:             time,
		FlightID:         flightID,
		FlightSnapshotID: flightSnapshotID,
		Cabin:            cabin,
		Resource:         resource,
		Class:            class,
		BookedSeats:      bookedSeats,
		NestingCapacity:  nestingCapacity,
		AvailableSeats:   availableSeats,
		TotalCash:        totalCash,
		LastBooking:      lastBooking,
		CorrectlyParsed:  correctlyParsed,
		Noshow:           noshow,
	}
}

func NewTicketSnapshotCreateFromTicketCreate(ticket *TicketCreate, flightSnapshotID oid.ID, t time.Time) *TicketSnapshotCreate {
	return NewTicketSnapshotCreate(
		t,
		ticket.FlightID,
		flightSnapshotID,
		ticket.Cabin,
		ticket.Resource,
		ticket.Class,
		ticket.BookedSeats,
		ticket.NestingCapacity,
		ticket.AvailableSeats,
		ticket.TotalCash,
		ticket.LastBooking,
		ticket.CorrectlyParsed,
		ticket.Noshow,
	)
}

func (tsC *TicketSnapshotCreate) ToEntity(id oid.ID) *entity.TicketSnapshot {
	return entity.NewTicketSnapshot(
		id,
		tsC.Time,
		tsC.FlightID,
		tsC.FlightSnapshotID,
		tsC.Cabin,
		tsC.Resource,
		tsC.Class,
		tsC.BookedSeats,
		tsC.NestingCapacity,
		tsC.AvailableSeats,
		tsC.TotalCash,
		tsC.LastBooking,
		tsC.CorrectlyParsed,
		tsC.Noshow,
	)
}
