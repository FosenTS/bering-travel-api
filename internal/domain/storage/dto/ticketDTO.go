package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"time"
)

type TicketCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	FlightID        oid.ID  `json:"flightID" db:"flight_id" binding:"required"`
	Cabin           string  `json:"cabin" db:"cabin" binding:"required"`
	Resource        int     `json:"resource" db:"resource" binding:"required"`
	Class           string  `json:"class" db:"class" binding:"required"`
	BookedSeats     int     `json:"bookedSeats" db:"booked_seats" binding:"required"`
	NestingCapacity int     `json:"nestingCapacity" db:"nesting_capacity" binding:"required"`
	AvailableSeats  int     `json:"availableSeats" db:"available_seats" binding:"required"`
	TotalCash       float64 `json:"totalCash" db:"total_cash" binding:"required"`

	LastBooking time.Time `json:"lastBooking" db:"last_booking" binding:"required"`

	LastUpdate time.Time `json:"relevance" db:"last_update" binding:"required"`

	CorrectlyParsed bool `json:"correct" db:"correctly_parsed" binding:"required"`

	Noshow           *int      `json:"noshow" db:"noshow" binding:"required"`
	NoshowLastUpdate time.Time `json:"noshowLastUpdate" db:"noshow_last_update" binding:"required"`
}

func NewTicketCreate(
	flightID oid.ID,
	cabin string,
	resource int,
	class string,
	bookedSeats int,
	nestingCapacity int,
	availableSeats int,
	totalCash float64,
	lastBooking time.Time,
	lastUpdate time.Time,
	correctlyParsed bool,
	noshow *int,
	noshowLastUpdate time.Time,
) *TicketCreate {
	return &TicketCreate{
		FlightID:         flightID,
		Cabin:            cabin,
		Resource:         resource,
		Class:            class,
		BookedSeats:      bookedSeats,
		NestingCapacity:  nestingCapacity,
		AvailableSeats:   availableSeats,
		TotalCash:        totalCash,
		LastBooking:      lastBooking,
		LastUpdate:       lastUpdate,
		CorrectlyParsed:  correctlyParsed,
		Noshow:           noshow,
		NoshowLastUpdate: noshowLastUpdate,
	}
}

func (tC *TicketCreate) ToEntity(id oid.ID) *entity.Ticket {
	return entity.NewTicket(
		id,
		tC.FlightID,
		tC.Cabin,
		tC.Resource,
		tC.Class,
		tC.BookedSeats,
		tC.NestingCapacity,
		tC.AvailableSeats,
		tC.TotalCash,
		tC.LastBooking,
		tC.LastUpdate,
		tC.CorrectlyParsed,
		tC.Noshow,
		tC.NoshowLastUpdate,
	)
}

type TicketSells struct {
	Class       string    `json:"class" db:"class" binding:"required"`
	BookedSeats int       `json:"bookedSeats" db:"booked_seats" binding:"required"`
	Time        time.Time `json:"time" db:"time"`
}

type TicketSellsSegment []TicketSells

type TicketIdentity struct {
	FlightID oid.ID
	Class    string
}

func NewTicketIdentity(flightID oid.ID, class string) *TicketIdentity {
	return &TicketIdentity{FlightID: flightID, Class: class}
}
