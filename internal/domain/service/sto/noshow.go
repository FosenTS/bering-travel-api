package sto

type FlightNoshow struct {
	F       FlightIdentity
	Classes map[string]uint
}
