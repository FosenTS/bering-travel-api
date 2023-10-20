package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
)

type MultiFlightCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	AirlCode string  `json:"airlCode" db:"airl_code" binding:"required"`
	Number   string  `json:"number" db:"number" binding:"required"`
	FLtDate  string  `json:"fltDate" db:"flt_date" binding:"required"`
	L13      *oid.ID `json:"l13" db:"l13" binding:"required"`
	L12      *oid.ID `json:"l12" db:"l12" binding:"required"`
	L23      *oid.ID `json:"l23" db:"l23" binding:"required"`
}

func NewMultiFlightCreate(
	airlCode string,
	number string,
	FLtDate string,
	l13 *oid.ID,
	l12 *oid.ID,
	l23 *oid.ID,
) *MultiFlightCreate {
	return &MultiFlightCreate{
		AirlCode: airlCode,
		Number:   number,
		FLtDate:  FLtDate,
		L13:      l13,
		L12:      l12,
		L23:      l23,
	}
}

func (m *MultiFlightCreate) ToEntity(id oid.ID) *entity.MultiFlight {
	return entity.NewMultiFlight(
		id,
		m.AirlCode,
		m.Number,
		m.FLtDate,
		m.L13,
		m.L12,
		m.L23,
	)
}
