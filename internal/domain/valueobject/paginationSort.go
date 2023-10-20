package valueobject

type PaginationSortBy string

type paginationSort struct{}

func (*paginationSort) Time() PaginationSortBy {
	return "Time"
}

func (*paginationSort) Load() PaginationSortBy {
	return "Load"
}

func (*paginationSort) DepartureTime() PaginationSortBy {
	return "DepartureTime"
}

func (*paginationSort) Origin() PaginationSortBy {
	return "Origin"
}

func (*paginationSort) Destination() PaginationSortBy {
	return "Destination"
}

func (*paginationSort) Number() PaginationSortBy {
	return "Class"
}

func (*paginationSort) MultiFlightOrder() PaginationSortBy {
	return "MultiFlightOrder"
}

func (ps *paginationSort) Contains(s string) bool {
	switch PaginationSortBy(s) {
	case ps.Time():
	case ps.Load():
	case ps.DepartureTime():
	case ps.Origin():
	case ps.Destination():
	case ps.Number():
	case ps.MultiFlightOrder():
	default:
		return false
	}

	return true
}

var PaginationSort = paginationSort{}
