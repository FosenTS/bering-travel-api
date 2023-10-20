package service

import (
	"bering-travel-api/internal/domain/storage"
)

type MapperService interface {
}

type mappperService struct {
	mapperStorage storage.MapperStorage
}

func NewMapperService(mapperStorage storage.MapperStorage) MapperService {
	return &mappperService{mapperStorage: mapperStorage}
}
