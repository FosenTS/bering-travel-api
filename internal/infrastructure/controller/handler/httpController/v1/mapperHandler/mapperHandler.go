package mapperHandler

import (
	"bering-travel-api/internal/domain/service"
	"bering-travel-api/internal/infrastructure/controller/safeObject"
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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
	r.GET("getPointers", h.getPointers)
}

func (h *Handler) addPointer(c *gin.Context) {

	var point *safeObject.Pointer
	err := c.BindJSON(&point)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, "bad request")
		return
	}

	err = h.mapperService.StorePointer(c, point)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "unkown error")
		return
	}

	c.JSON(http.StatusOK, "pointer added")
}

func (h *Handler) getPointers(c *gin.Context) {

	pointers, err := h.mapperService.GetAllPointes(c)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "unkown error")
		return
	}

	c.JSON(http.StatusOK, pointers)

}
