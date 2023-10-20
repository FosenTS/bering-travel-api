package sto

import "saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"

type FlightSegment struct {
	Flight  *entity.Flight
	Tickets []*entity.Ticket
}
