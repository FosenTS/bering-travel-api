package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"time"
)

type RivalFileCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	Name            string    `json:"name" db:"name" binding:"required"`
	NumBytes        int64     `json:"numBytes" db:"num_bytes" binding:"required"`
	ParsedAt        time.Time `json:"parsedAt" db:"parsed_at" binding:"required"`
	CorrectlyParsed bool      `json:"correctlyParsed" db:"correctly_parsed" binding:"required"`
}

func NewRivalFlightsFileCreate(name string, numBytes int64, parsedAt time.Time, correctlyParsed bool) *RivalFileCreate {
	return &RivalFileCreate{Name: name, NumBytes: numBytes, ParsedAt: parsedAt, CorrectlyParsed: correctlyParsed}
}

func (rffC *RivalFileCreate) ToEntity(id oid.ID) *entity.RivalFile {
	return entity.NewRivalFlightsFile(id, rffC.Name, rffC.NumBytes, rffC.ParsedAt, rffC.CorrectlyParsed)
}
