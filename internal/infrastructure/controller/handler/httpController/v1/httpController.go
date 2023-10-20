package httpHandler

import (
	"bering-travel-api/internal/application/config"
	"github.com/gin-gonic/gin"
)

type HTTPHandler interface {
	RegisterRouter(group *gin.RouterGroup)
}

type httpHandlerController struct {
	cfg           config.HandlerConfig
	mapperHandler HTTPHandler
}

func NewHTTPHandlerController(
	cfg config.HandlerConfig,
	mapperHandler HTTPHandler,
) HTTPHandler {
	return &httpHandlerController{
		cfg: cfg,

		mapperHandler: mapperHandler,
	}
}

func (h *httpHandlerController) RegisterRouter(r *gin.RouterGroup) {
	api := r.Group("/api")

	h.mapperHandler.RegisterRouter(api)
}
