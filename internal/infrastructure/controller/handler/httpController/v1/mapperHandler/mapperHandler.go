package mapperHandler

import (
	"bering-travel-api/internal/domain/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	mapperService service.MapperService
}

func NewHandler(mapperService service.MapperService) *Handler {
	return &Handler{mapperService: mapperService}
}

func (h *Handler) RegisterRouter(r *gin.RouterGroup) {
	r.GET("")
}
