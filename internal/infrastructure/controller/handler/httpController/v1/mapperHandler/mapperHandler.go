package mapperHandler

import (
	"bering-travel-api/internal/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	mapperService service.MapperService
}

func NewHandler(mapperService service.MapperService) *Handler {
	return &Handler{mapperService: mapperService}
}

func (h *Handler) RegisterRouter(r *gin.RouterGroup) {
	r.POST("addPointer", h.addPointer)
}

func (h *Handler) addPointer(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
