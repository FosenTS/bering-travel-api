package sto

import "saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"

type RivalHead struct {
	LowestCostRival  *entity.RivalFlight `json:"lowestCost" db:"lowest_cost" binding:"required"`
	NearestTimeRival *entity.RivalFlight `json:"nearestTimeRival" db:"nearest_time_rival" binding:"required"`
}

func NewRivalHead(lowestCostRival *entity.RivalFlight, nearestTimeRival *entity.RivalFlight) *RivalHead {
	return &RivalHead{LowestCostRival: lowestCostRival, NearestTimeRival: nearestTimeRival}
}
