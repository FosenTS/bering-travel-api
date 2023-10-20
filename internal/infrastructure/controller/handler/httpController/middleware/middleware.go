package middleware

import (
	"fmt"
	"github.com/achillescres/pkg/security/ajwt"
	"github.com/gin-gonic/gin"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/application/config"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/service"
)

var (
	errEmptyHeader = fmt.Errorf(
		"error auth header is empty",
	)
	errInvalidHeader = fmt.Errorf(
		"error invalid auth header",
	)
)

const (
	userLoginKey        = "userLoginKey"
	userKey             = "userKey"
	authorizationHeader = "Authorization"
	bearer              = "Bearer"
)

//go:generate go run github.com/vektra/mockery/v2@v2.33.2 --name=Middleware
type Middleware interface {
	UserPolicy(c *gin.Context)
	Jwt(c *gin.Context) error
	FlightIdentityFromQueries(c *gin.Context)
	SortingMiddleware(c *gin.Context)
}

type middleware struct {
	jwtManager       ajwt.JWTManager
	policyService    service.PolicyService
	middlewareConfig config.MiddlewareConfig
}

func NewMiddleware(jwtManager ajwt.JWTManager, middlewareConfig config.MiddlewareConfig, authService service.PolicyService) Middleware {
	return &middleware{jwtManager: jwtManager, middlewareConfig: middlewareConfig, policyService: authService}
}
