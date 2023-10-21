package service

import (
	"bering-travel-api/internal/domain/entity"
	"bering-travel-api/internal/domain/storage"
	"bering-travel-api/internal/infrastructure/controller/safeObject"
	"context"
)

type MapperService interface {
	StorePointer(ctx context.Context, pointer *safeObject.Pointer) error
	GetAllPointes(ctx context.Context) ([]*entity.Pointer, error)
}

type mappperService struct {
	mapperStorage storage.MapperStorage
}

func NewMapperService(mapperStorage storage.MapperStorage) MapperService {
	return &mappperService{mapperStorage: mapperStorage}
}

func (m *mappperService) StorePointer(ctx context.Context, pointer *safeObject.Pointer) error {
	err := m.mapperStorage.StorePointer(ctx, pointer)
	if err != nil {
		return err
	}
	return nil
}

func (m *mappperService) GetAllPointes(ctx context.Context) ([]*entity.Pointer, error) {
	pointers, err := m.mapperStorage.GetAllPointers(ctx)
	if err != nil {
		return nil, err
	}

	return pointers, nil
}
