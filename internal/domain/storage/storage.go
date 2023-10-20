package storage

import (
	"bering-travel-api/internal/domain/entity"
	"bering-travel-api/internal/domain/storage/dto"
	"context"
	"fmt"
	"github.com/FosenTS/pkg/object/oid"
	"github.com/Masterminds/squirrel"
)

type Storage[Entity entity.Entity, Create dto.Create] interface {
	GetByID(ctx context.Context, id oid.ID) (*Entity, error)
	GetAll(ctx context.Context) ([]*Entity, error)
	GetAllInMap(ctx context.Context) (map[oid.ID]*Entity, error)
	Store(ctx context.Context, c *Create) (*Entity, error)
	StoreMany(ctx context.Context, cs []*Create) ([]*Entity, error)
	UpdateByID(ctx context.Context, f *Entity) error
	DeleteByID(ctx context.Context, id oid.ID) (*Entity, error)
}

type Querier interface {
	ToSql() (string, []interface{}, error)
}

type QueryBuilder[Builder Querier] interface {
	Querier
	OrderBy(orderBys ...string) Builder
	Where(pred interface{}, args ...interface{}) Builder
	Limit(limit uint64) Builder
}

type Option[Builder Querier] func(QueryBuilder[Builder])

func SortOption[Builder Querier](by, order string) Option[Builder] {
	return func(builder QueryBuilder[Builder]) {
		builder.OrderBy(fmt.Sprintf("%s %s", by, order))
	}
}

func GtOrEqOption[Builder Querier](by, value string) Option[Builder] {
	return func(builder QueryBuilder[Builder]) {
		builder.Where(squirrel.GtOrEq{by: value})
	}
}

func GtOption[Builder Querier](by, value string) Option[Builder] {
	return func(builder QueryBuilder[Builder]) {
		builder.Where(squirrel.Gt{by: value})
	}
}

func LimitOption[Builder Querier](limit uint64) Option[Builder] {
	return func(builder QueryBuilder[Builder]) {
		builder.Limit(limit)
	}
}

func SortHeadOption[Builder Querier](by, order string, limit uint64) Option[Builder] {
	return func(builder QueryBuilder[Builder]) {
		SortOption[Builder](by, order)(builder)
		LimitOption[Builder](limit)(builder)
	}
}
