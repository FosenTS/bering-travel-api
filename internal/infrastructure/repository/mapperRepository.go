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
		"INSERT INTO pointer (name, description, latitude, longitude, rating) VALUES ($1, $2, $3, $4, $5)",
		pointer.Name,
		pointer.Description,
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
		"SELECT id, name, description, latitude, longitude, rating FROM pointer",
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
		"SELECT id, name, description, latitude, longitude, rating FROM pointer WHERE id = $1",
		id,
	).Scan(
		&pointer.Id,
		&pointer.Name,
		&pointer.Description,
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
		"INSERT INTO user_visit (user_id, pointer_id, rating, comment, user_activity) VALUES ($1, $2, $3, $4, $5)",
		visit.UserId,
		visit.PointerId,
		visit.Rating,
		visit.Comment,
		visit.Activity,
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

func (m *mapperRepository) GetCoins(ctx context.Context, id oid.ID) (uint, error) {
	var coins uint
	err := m.pool.QueryRow(
		ctx,
		"SELECT coin_count FROM user_coins WHERE user_id=$1",
		id,
	).Scan(&coins)
	if err != nil {
		return 0, err
	}

	return coins, nil
}

func (m *mapperRepository) UpdateCoins(ctx context.Context, id oid.ID, count int) error {
	var countRow int
	err := m.pool.QueryRow(
		ctx,
		"SELECT count(*) FROM user_coins WHERE user_id=$1",
		id,
	).Scan(&countRow)
	if err != nil {
		return err
	}
	if countRow == 0 {
		_, err := m.pool.Exec(
			ctx,
			"INSERT INTO user_coins(user_id, coin_count) VALUES ($1, $2)",
			id,
			count,
		)
		if err != nil {
			return err
		}
		return nil
	}

	_, err = m.pool.Exec(
		ctx,
		"UPDATE user_coins SET coin_count=coin_count+$2 WHERE user_id = $1",
		id,
		count,
	)

	if err != nil {
		return err
	}
	return nil
}

func (m *mapperRepository) GetCommentsByPlace(ctx context.Context, point_id uint) ([]*entity.Commit, error) {
	rows, err := m.pool.Query(
		ctx,
		"Select user_id, comment FROM user_visit WHERE pointer_id=$1",
		point_id,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	comments := make([]*entity.Commit, 0)
	err = pgxscan.ScanAll(&comments, rows)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
