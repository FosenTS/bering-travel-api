package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
)

type RateCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	AirlCode        string `json:"airlCode" db:"airl_code" binding:"required"`
	OriginIATA      string `json:"origIATA" db:"origin_iata" binding:"required"`
	DestinationIATA string `json:"destIATA" db:"destination_iata" binding:"required"`
	Class           string `json:"class" db:"class" binding:"required"`
	SubClass        string `json:"subClass" db:"subClass" binding:"required"`
	JourneyStart    string `json:"journeyStart" db:"journey_start" binding:"required"`
	JourneyStop     string `json:"journeyStop" db:"journey_stop" binding:"required"`
	SaleStart       string `json:"saleStart" db:"sale_start" binding:"required"`
	SaleStop        string `json:"saleStop" db:"sale_stop" binding:"required"`
	CurrentCash     uint   `json:"currentCash" db:"current_cash" binding:"required"`
}

func NewRateCreate(airlCode string, originIATA string, destinationIATA string, class string, subClass string, journeyStart string, journeyStop string, saleStart string, saleStop string, currentCash uint) *RateCreate {
	return &RateCreate{AirlCode: airlCode, OriginIATA: originIATA, DestinationIATA: destinationIATA, Class: class, SubClass: subClass, JourneyStart: journeyStart, JourneyStop: journeyStop, SaleStart: saleStart, SaleStop: saleStop, CurrentCash: currentCash}
}

func (rC *RateCreate) ToEntityWithFilledAdditionalFields(id oid.ID) *entity.Rate {
	return entity.NewRate(
		id,
		rC.AirlCode,
		rC.OriginIATA,
		rC.DestinationIATA,
		rC.Class,
		rC.SubClass,
		rC.JourneyStart,
		rC.JourneyStop,
		rC.SaleStart,
		rC.SaleStop,
		rC.CurrentCash,
	)
}
