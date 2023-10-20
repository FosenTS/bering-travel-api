package repository

import (
	"bering-travel-api/internal/domain/storage"
	"github.com/FosenTS/pkg/db/postgresql"
)

type MapperRepository storage.MapperStorage

type mapperRepository struct {
	pool postgresql.PGXPool
}

func NewMapperRepository(pool postgresql.PGXPool) MapperRepository {
	return &mapperRepository{pool: pool}
}
