package repository

import (
	"bering-travel-api/internal/domain/entity"
	"bering-travel-api/internal/domain/storage"
	"bering-travel-api/internal/infrastructure/controller/safeObject"
	"context"
	"errors"
	"github.com/achillescres/pkg/db/postgresql"
	"github.com/achillescres/pkg/object/oid"
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

func (m *mapperRepository) GetPointerById(ctx context.Context, id uint) (*entity.Pointer, error) {
	var pointer entity.Pointer
	err := m.pool.QueryRow(
		ctx,
		"SELECT id, name, address, latitude, longitude, rating FROM pointer WHERE id = $1",
		id,
	).Scan(
		&pointer.Id,
		&pointer.Name,
		&pointer.Address,
		&pointer.Latitude,
		&pointer.Longitude,
		&pointer.Rating,
	)
	if err != nil {
		return nil, err
	}

	return &pointer, err
}

func (m *mapperRepository) StoreUserVisit(ctx context.Context, visit *safeObject.UserVisit) error {
	_, err := m.pool.Exec(
		ctx,
		"INSERT INTO user_visit (user_id, pointer_id, rating, comment) VALUES ($1, $2, $3, $4)",
		visit.UserId,
		visit.PointerId,
		visit.Rating,
		visit.Comment,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *mapperRepository) GetUserVisitById(ctx context.Context, id oid.ID) ([]*entity.UserVisit, error) {
	rows, err := m.pool.Query(
		ctx,
		"SELECT id, user_id, pointer_id, rating, comment from user_visit where user_id=$1",
		id,
	)
	if err != nil {
		return nil, err
	}

	visits := make([]*entity.UserVisit, 0)
	err = pgxscan.ScanAll(&visits, rows)
	if err != nil {
		return nil, err
	}
	return visits, nil
}
func (m *mapperRepository) GetUserVisitCount(ctx context.Context, id oid.ID) (uint, error) {
	var count uint
	err := m.pool.QueryRow(
		ctx,
		"SELECT count(*) FROM user_visit where user_id=$1",
		id,
	).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
