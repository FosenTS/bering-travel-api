package sto

import "time"

type BookingDataChange struct {
	AirlCode        string
	Number          string
	OriginIATA      string
	DestinationIATA string
	FltDate         string
	Class           string
	NestingCapacity uint
	Time            time.Time
	UserLogin       string
	Exported        bool
}

func NewBookingDataChange(airlCode string, number string, originIATA string, destinationIATA string, fltDate string, class string, nestingCapacity uint, time time.Time, userLogin string, exported bool) *BookingDataChange {
	return &BookingDataChange{AirlCode: airlCode, Number: number, OriginIATA: originIATA, DestinationIATA: destinationIATA, FltDate: fltDate, Class: class, NestingCapacity: nestingCapacity, Time: time, UserLogin: userLogin, Exported: exported}
}

type BookingDataChangeSegment []BookingDataChange
