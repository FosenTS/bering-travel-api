package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"time"
)

type BookingDataChangeCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	SegmentUUID     string    `json:"segmentUUID" db:"segment_uuid" binding:"required"`
	AirlCode        string    `json:"airlCode" db:"airl_code" binding:"required,alphanum"`
	Number          string    `json:"fltNum" db:"number" binding:"required"`
	OriginIATA      string    `json:"originIATA" db:"orig_iata" binding:"required,alphanum"`
	DestinationIATA string    `json:"destinationIATA" db:"dest_iata" binding:"required,alphanum"`
	FltDate         string    `json:"date" db:"flt_date" binding:"required,len=8" time_format:"20060102"`
	Class           string    `json:"class" db:"class" binding:"required,len=1,alphanum"`
	NestingCapacity uint      `json:"nestingCapacity" db:"nesting_capacity" binding:"required"`
	Time            time.Time `json:"time" db:"time" binding:"required"`
	UserLogin       string    `json:"userLogin" db:"user_login" binding:"required"`
	Exported        bool      `json:"exported" db:"exported" binding:"required"`
}

func NewBookingDataChangeCreate(segmentUUID string, airlCode string, number string, originIATA string, destinationIATA string, fltDate string, class string, nestingCapacity uint, time time.Time, userLogin string, exported bool) *BookingDataChangeCreate {
	return &BookingDataChangeCreate{SegmentUUID: segmentUUID, AirlCode: airlCode, Number: number, OriginIATA: originIATA, DestinationIATA: destinationIATA, FltDate: fltDate, Class: class, NestingCapacity: nestingCapacity, Time: time, UserLogin: userLogin, Exported: exported}
}

func (c *BookingDataChangeCreate) ToEntity(id oid.ID, lastNestingCapacity uint) *entity.BookingDataChange {
	return entity.NewBookingDataChange(
		id,
		c.SegmentUUID,
		c.AirlCode,
		c.Number,
		c.OriginIATA,
		c.DestinationIATA,
		c.FltDate,
		c.Class,
		c.NestingCapacity,
		c.Time,
		c.UserLogin,
		c.Exported,
		lastNestingCapacity,
	)
}
