package sto

import "github.com/achillescres/pkg/object/oid"

type FlightComment struct {
	FlightID oid.ID
	Text     string
}

func NewFlightComment(flightID oid.ID, text string) *FlightComment {
	return &FlightComment{FlightID: flightID, Text: text}
}
