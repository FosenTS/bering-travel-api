package storage

import (
	"bering-travel-api/internal/domain/entity"
	"bering-travel-api/internal/infrastructure/controller/safeObject"
	"context"
	"github.com/achillescres/pkg/object/oid"
)

type MapperStorage interface {
	StorePointer(ctx context.Context, pointer *safeObject.Pointer) error
	GetAllPointers(ctx context.Context) ([]*entity.Pointer, error)
	GetPointerById(ctx context.Context, id uint) (*entity.Pointer, error)
	StoreUserVisit(ctx context.Context, visit *safeObject.UserVisit) error
	GetUserVisitById(ctx context.Context, id oid.ID) ([]*entity.UserVisit, error)
	GetCommentsByPlace(ctx context.Context, point_id uint) ([]*entity.Commit, error)
	GetAllRatingUsers(ctx context.Context, point_id uint) ([]uint, error)
}
