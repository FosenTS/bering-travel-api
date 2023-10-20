package sto

import "saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/storage/dto"

type FlightIdentity struct {
	AirlCode        string `json:"airlCode" form:"airlCode" binding:"required"`
	Number          string `json:"number" form:"number" binding:"required"`
	OriginIATA      string `json:"originIATA" form:"originIATA" binding:"required"`
	DestinationIATA string `json:"destinationIATA" form:"destinationIATA" binding:"required"`
	FltDate         string `json:"fltDate" form:"fltDate" binding:"required"`
}

func NewFlightIdentity(airlCode string, number string, originIATA string, destinationIATA string, fltDate string) *FlightIdentity {
	return &FlightIdentity{AirlCode: airlCode, Number: number, OriginIATA: originIATA, DestinationIATA: destinationIATA, FltDate: fltDate}
}

func (fi *FlightIdentity) ToDTO() *dto.FlightIdentity {
	return dto.NewFlightIdentity(
		fi.AirlCode,
		fi.Number,
		fi.OriginIATA,
		fi.DestinationIATA,
		fi.FltDate,
	)
}
