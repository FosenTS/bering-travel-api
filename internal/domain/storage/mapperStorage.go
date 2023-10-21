package storage

import (
	"bering-travel-api/internal/domain/entity"
	"bering-travel-api/internal/infrastructure/controller/safeObject"
	"context"
)

type MapperStorage interface {
	StorePointer(ctx context.Context, pointer *safeObject.Pointer) error
	GetAllPointers(ctx context.Context) ([]*entity.Pointer, error)
}
