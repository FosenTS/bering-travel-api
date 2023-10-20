package product

import (
	"bering-travel-api/internal/domain/storage"
	"bering-travel-api/internal/infrastructure/repository"
	"github.com/FosenTS/pkg/db/postgresql"
)

type Storages struct {
	MapperStorage storage.MapperStorage
}

func NewStorages(pgPool postgresql.PGXPool) *Storages {
	return &Storages{
		MapperStorage: repository.NewMapperRepository(pgPool),
	}
}
