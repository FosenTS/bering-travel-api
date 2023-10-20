package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
)

type FlightTagCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	Text string `json:"text" db:"text" binding:"required"`
}

func NewFlightTagCreate(text string) *FlightTagCreate {
	return &FlightTagCreate{Text: text}
}

func (c *FlightTagCreate) ToEntity(id oid.ID) *entity.FlightTag {
	return entity.NewFlightTag(
		id,
		c.Text,
	)
}
