package middleware

import (
	"context"
	"errors"
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/valueobject"
)

const restrictionsKey = "restrictions"

func (m *middleware) UserPolicy(c *gin.Context) {
	err := m.Jwt(c)
	if err != nil {
		log.Errorln(err)
		return
	}

	user, err := GetUser(c)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "couldn't get user")
		return
	}

	c.Set(restrictionsKey, *entity.NewRestrictions(user.AirlineCode, valueobject.UserRoles[user.Role]))
	c.Next()
}

func GetRestrictions(ctx context.Context) (entity.Restrictions, error) {
	val, ok := ctx.Value(restrictionsKey).(entity.Restrictions)
	if !ok {
		return entity.Restrictions{}, errors.New("error no restrictions was found")
	}
	return val, nil
}

func SetRestrictionForDeveloper(ctx context.Context) context.Context {
	return context.WithValue(ctx, restrictionsKey, *entity.NewRestrictions("", valueobject.UserRoleDeveloper))
}
