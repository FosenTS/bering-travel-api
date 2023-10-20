package sto

type RouteUpdate struct {
	Distance             float64
	TemporaryConsumption float64
	ConstantConsumption  float64
}

func NewRouteUpdate(distance float64, temporaryConsumption float64, constantConsumption float64) *RouteUpdate {
	return &RouteUpdate{Distance: distance, TemporaryConsumption: temporaryConsumption, ConstantConsumption: constantConsumption}
}
