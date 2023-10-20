package sto

import (
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
)

type FlightTable struct {
	Flight                 *entity.Flight           `json:"flight" binding:"required"`
	Snapshots              []*entity.FlightSnapshot `json:"snapshots" binding:"required"`
	SimilarWeekAgoSnapshot *entity.FlightSnapshot   `json:"similarWeekAgoSnapshot" binding:"required"`
	Tickets                []*entity.Ticket         `json:"tickets" binding:"required"`
}
