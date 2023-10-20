package dto

type FlightIdentity struct {
	AirlCode        string `json:"airlCode" binding:"required"`
	Number          string `json:"number" binding:"required"`
	OriginIATA      string `json:"originIATA" binding:"required"`
	DestinationIATA string `json:"destinationIATA" binding:"required"`
	FltDate         string `json:"fltDate" binding:"required"`
}

func NewFlightIdentity(airlCode string, number string, originIATA string, destinationIATA string, fltDate string) *FlightIdentity {
	return &FlightIdentity{AirlCode: airlCode, Number: number, OriginIATA: originIATA, DestinationIATA: destinationIATA, FltDate: fltDate}
}
