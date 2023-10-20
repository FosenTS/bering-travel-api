package product

import (
	"bering-travel-api/internal/application/config"
	httpHandler "bering-travel-api/internal/infrastructure/controller/handler/httpController/v1"
	"bering-travel-api/internal/infrastructure/controller/handler/httpController/v1/mapperHandler"
	"github.com/gin-gonic/gin"
)

type Controllers struct {
	handler httpHandler.HTTPHandler
}

func NewControllers(
	gateways *Gateways,
	handlerCfg config.HandlerConfig,
) (*Controllers, error) {
	return &Controllers{
		handler: httpHandler.NewHTTPHandlerController(
			handlerCfg,
			mapperHandler.NewHandler(gateways.Services.MapperService),
		),
	}, nil
}

func (c *Controllers) RegisterHandlers(r *gin.RouterGroup) {
	c.handler.RegisterRouter(r)
}
