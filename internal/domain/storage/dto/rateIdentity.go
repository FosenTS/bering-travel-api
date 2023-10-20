package dto

type RateIdentity struct {
	AirlCode        string `json:"airlCode" db:"airl_code" binding:"required"`
	OriginIATA      string `json:"origIATA" db:"origin_iata" binding:"required"`
	DestinationIATA string `json:"destIATA" db:"destination_iata" binding:"required"`
	Class           string `json:"class" db:"class" binding:"required"`
	SubClass        string `json:"subClass" db:"subClass" binding:"required"`
	JourneyStart    string `json:"journeyStart" db:"journey_start" binding:"required"`
	JourneyStop     string `json:"journeyStop" db:"journey_stop" binding:"required"`
	SaleStart       string `json:"saleStart" db:"sale_start" binding:"required"`
	SaleStop        string `json:"saleStop" db:"sale_stop" binding:"required"`
}

func NewRateIdentity(airlCode string, originIATA string, destinationIATA string, class string, subClass string, journeyStart string, journeyStop string, saleStart string, saleStop string) *RateIdentity {
	return &RateIdentity{AirlCode: airlCode, OriginIATA: originIATA, DestinationIATA: destinationIATA, Class: class, SubClass: class, JourneyStart: journeyStart, JourneyStop: journeyStop, SaleStart: saleStart, SaleStop: saleStop}
}
