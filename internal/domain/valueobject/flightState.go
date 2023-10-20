package valueobject

type flightState struct {
	Iop string
}

func (*flightState) Open() string {
	return "open"
}

func (*flightState) Closed() string {
	return "closed"
}

func (*flightState) Changed() string {
	return "changed"
}

func (*flightState) Removed() string {
	return "removed"
}

var FlightState = flightState{}
