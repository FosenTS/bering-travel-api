package product

import (
	"bering-travel-api/internal/domain/service"
)

type Services struct {
	MapperService service.MapperService
}

func NewServices(
	repos *Storages,
) (*Services, error) {
	return &Services{
		MapperService: service.NewMapperService(repos.MapperStorage),
	}, nil
}
