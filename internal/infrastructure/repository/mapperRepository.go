package repository

import (
	"bering-travel-api/internal/domain/entity"
	"bering-travel-api/internal/domain/storage"
	"bering-travel-api/internal/infrastructure/controller/safeObject"
	"context"
	"errors"
	"github.com/achillescres/pkg/db/postgresql"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

type MapperRepository storage.MapperStorage

type mapperRepository struct {
	pool postgresql.PGXPool
}

func NewMapperRepository(pool postgresql.PGXPool) MapperRepository {
	return &mapperRepository{pool: pool}
}

func (m *mapperRepository) StorePointer(ctx context.Context, pointer *safeObject.Pointer) error {
	_, err := m.pool.Exec(
		ctx,
		"INSERT INTO pointer (name, address, latitude, longitude, rating) VALUES ($1, $2, $3, $4, $5)",
		pointer.Name,
		pointer.Address,
		pointer.Latitude,
		pointer.Longitude,
		pointer.Rating,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m *mapperRepository) GetAllPointers(ctx context.Context) ([]*entity.Pointer, error) {
	rows, err := m.pool.Query(
		ctx,
		"SELECT id, name, address, latitude, longitude, rating FROM pointer",
	)
	if err != nil {
		return nil, err
	}

	pointers := make([]*entity.Pointer, 0)
	err = pgxscan.ScanAll(&pointers, rows)
	if errors.Is(err, pgx.ErrNoRows) {
		return []*entity.Pointer{}, nil
	}

	if err != nil {
		return nil, err
	}

	return pointers, nil
}
