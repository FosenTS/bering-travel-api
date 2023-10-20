package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"time"
)

type BookingDataFileCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	Name            string    `json:"name" db:"name" binding:"required"`
	NumBytes        int64     `json:"num_bytes" db:"num_bytes" binding:"required"`
	Time            time.Time `json:"time" db:"time" binding:"required"`
	ParsedAt        time.Time `json:"parsed_at" db:"parsed_at" binding:"required"`
	CorrectlyParsed bool      `json:"correctlyParsed" db:"correctly_parsed" binding:"required"`
}

func NewBookingDataFileCreate(name string, numBytes int64, time time.Time, parsedAt time.Time, correctlyParsed bool) *BookingDataFileCreate {
	return &BookingDataFileCreate{Name: name, NumBytes: numBytes, Time: time, ParsedAt: parsedAt, CorrectlyParsed: correctlyParsed}
}

func (bfC *BookingDataFileCreate) ToEntity(id oid.ID) *entity.BookingDataFile {
	return entity.NewBookingDataFile(id, bfC.Name, bfC.NumBytes, bfC.Time, bfC.ParsedAt, bfC.CorrectlyParsed)
}
