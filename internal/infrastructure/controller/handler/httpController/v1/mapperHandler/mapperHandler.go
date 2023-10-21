package mapperHandler

import (
	"bering-travel-api/internal/domain/service"
	"bering-travel-api/internal/infrastructure/controller/safeObject"
	"errors"
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/achillescres/pkg/object/oid"
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
	r.POST("addUserVisit", h.addUserVisit)
	r.GET("getUserVisits", h.getUserVisits)
	r.GET("getPlaces", h.getPlaces)
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

func (h *Handler) addUserVisit(c *gin.Context) {

	var userVisits *safeObject.UserVisit
	err := c.BindJSON(&userVisits)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, "bad request")
		return
	}
	err = h.mapperService.StoreUserVisit(c, userVisits)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "unkown error")
		return
	}

	c.JSON(http.StatusOK, "success")
}

func (h *Handler) getUserVisits(c *gin.Context) {

	userId, ok := c.GetQuery("id")
	if !ok {
		err := errors.New("empty id query field")
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, "bad request")
		return
	}

	id := oid.ToID(userId)

	visits, err := h.mapperService.GetUserVisitsById(c, id)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "unknown error")
		return
	}

	c.JSON(http.StatusOK, visits)
}

func (h *Handler) getPlaces(c *gin.Context) {

	sortType, ok := c.GetQuery("sort")
	if !ok {
		err := errors.New("empty sort query")
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnprocessableEntity, err, "empty sort query")
		return
	}

	places, err := h.mapperService.GetPlaces(c, sortType)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "unkown error")
		return
	}

	c.JSON(http.StatusOK, places)
}
