package product

type Gateways struct {
	Services *Services
}

func NewGateways(services *Services) (*Gateways, error) {
	return &Gateways{
		Services: services,
	}, nil
}
