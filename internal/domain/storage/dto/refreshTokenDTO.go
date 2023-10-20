package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
)

type RefreshTokenCreate struct {
	Create             `json:"-" db:"-" binding:"-"`
	Token              string `json:"token" db:"login" binding:"required"`
	Login              string `json:"login" db:"login" binding:"required"`
	ExpirationTimeUnix int64  `json:"expirationTime" db:"expiration_time_unix" binding:"required"`
	CreateTimeUnix     int64  `json:"createTimeUnix" db:"create_time_unix" binding:"required"`
}

func NewRefreshTokenCreate(token string, login string, expirationTimeUnix int64, createTimeUnix int64) *RefreshTokenCreate {
	return &RefreshTokenCreate{Token: token, Login: login, ExpirationTimeUnix: expirationTimeUnix, CreateTimeUnix: createTimeUnix}
}

func (rTC RefreshTokenCreate) ToEntity(id oid.ID) *entity.RefreshToken {
	return entity.NewRefreshToken(id, rTC.Token, rTC.ExpirationTimeUnix, rTC.CreateTimeUnix)
}
