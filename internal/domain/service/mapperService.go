package service

import (
	"bering-travel-api/internal/domain/entity"
	"bering-travel-api/internal/domain/storage"
	"bering-travel-api/internal/infrastructure/controller/safeObject"
	"context"
	"github.com/achillescres/pkg/object/oid"
	"sort"
)

type MapperService interface {
	StorePointer(ctx context.Context, pointer *safeObject.Pointer) error
	GetAllPointes(ctx context.Context) ([]*entity.Pointer, error)
	StoreUserVisit(ctx context.Context, visit *safeObject.UserVisit) error
	GetUserVisitsById(ctx context.Context, id oid.ID) ([]*entity.UserVisit, error)
	GetPlaces(ctx context.Context, sortType string) ([]*entity.Places, error)
	GetUserProfile(ctx context.Context, id oid.ID) (*entity.Profile, error)
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

	for i, p := range pointers {
		ratings, err := m.mapperStorage.GetAllRatingUsers(ctx, p.Id)
		if err != nil {
			return nil, err
		}
		if ratings == nil {
			continue
		}
		var total uint
		ratings = append(ratings, p.Rating)
		for _, number := range ratings {
			total = total + number
		}
		pointers[i].Rating = uint(int(total) / len(ratings))
	}

	return pointers, nil
}

func (m *mappperService) StoreUserVisit(ctx context.Context, visit *safeObject.UserVisit) error {
	err := m.mapperStorage.StoreUserVisit(ctx, visit)
	if err != nil {
		return err
	}
	return nil
}

func (m *mappperService) GetUserVisitsById(ctx context.Context, id oid.ID) ([]*entity.UserVisit, error) {
	visits, err := m.mapperStorage.GetUserVisitById(ctx, id)
	if err != nil {
		return nil, err
	}

	for i, v := range visits {
		point, err := m.mapperStorage.GetPointerById(ctx, v.PointerId)
		if err != nil {
			return nil, err
		}

		visits[i].Pointer = *point
	}
	return visits, nil
}

func (m *mappperService) GetUserProfile(ctx context.Context, id oid.ID) (*entity.Profile, error) {
	visits, err := m.mapperStorage.GetUserVisitById(ctx, id)
	if err != nil {
		return nil, err
	}

	countActivity := 0
	for _, v := range visits {
		if v.Activity == true {
			countActivity++
		}
	}

	percent := countActivity / len(visits)

	return &entity.Profile{
		UserId:     id,
		VisitCount: uint(len(visits)),
		UserRating: uint(percent * 10),
	}, err
}

func (m *mappperService) GetPlaces(ctx context.Context, sortType string) ([]*entity.Places, error) {
	pointers, err := m.mapperStorage.GetAllPointers(ctx)
	if err != nil {
		return nil, err
	}

	for i, p := range pointers {
		ratings, err := m.mapperStorage.GetAllRatingUsers(ctx, p.Id)
		if err != nil {
			return nil, err
		}
		if ratings == nil {
			continue
		}
		var total uint
		ratings = append(ratings, p.Rating)
		for _, number := range ratings {
			total = total + number
		}
		pointers[i].Rating = uint(int(total) / len(ratings))
	}

	places := make([]*entity.Places, 0)

	for _, p := range pointers {
		commentaties, err := m.mapperStorage.GetCommentsByPlace(ctx, p.Id)
		if err != nil {
			return nil, err
		}

		sort.Slice(commentaties, func(i, j int) bool {
			return commentaties[i].Time.Before(commentaties[j].Time)
		})

		places = append(
			places,
			&entity.Places{
				Id:          p.Id,
				Name:        p.Name,
				Description: p.Description,
				Rating:      p.Rating,
				Latitude:    p.Latitude,
				Longitude:   p.Longitude,
				Time:        p.Time,
				Comments:    commentaties,
			},
		)

		switch sortType {
		case "rating":
			sort.Slice(places, func(i, j int) bool {
				return places[i].Rating > places[j].Rating
			})
		case "comments":
			sort.Slice(places, func(i, j int) bool {
				return len(places[i].Comments) > len(places[j].Comments)
			})
		case "time":
			sort.Slice(places, func(i, j int) bool {
				return places[j].Time.Before(places[i].Time)
			})
		}

	}
	return places, nil
}
