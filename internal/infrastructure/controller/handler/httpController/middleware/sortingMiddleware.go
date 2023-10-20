package middleware

import (
	"fmt"
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/valueobject"
	"strconv"
)

func (m *middleware) SortingMiddleware(c *gin.Context) {
	var sortBy valueobject.PaginationSortBy
	sortByS, ok := c.GetQuery("sortBy")
	if valueobject.PaginationSort.Contains(sortByS) {
		sortBy = valueobject.PaginationSortBy(sortByS)
	} else {
		err := fmt.Errorf("error sortBy query is invalid: %s", sortByS)
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusBadRequest, err, err.Error())
		return
	}

	var loadS string
	var load int
	var origin string
	var destination string
	var number string

	var err error

	switch valueobject.PaginationSortBy(sortByS) {
	case valueobject.PaginationSort.Time():
	case valueobject.PaginationSort.Load():
		loadS, ok = c.GetQuery("load")
		if !ok {
			log.Errorln(ginresponse.MissingRequiredQuery(c, "load"))
			return
		} else {
			load, err = strconv.Atoi(loadS)
			if err != nil {
				log.Errorln(err)
				ginresponse.ErrorString(c, http.StatusBadRequest, err, "error query parameter load is invalid")
				return
			}
		}
	case valueobject.PaginationSort.Origin():
		origin, ok = c.GetQuery("origin")
		if !ok {
			log.Errorln(ginresponse.MissingRequiredQuery(c, "origin"))
			return
		}
	case valueobject.PaginationSort.Destination():
		destination, ok = c.GetQuery("destination")
		if !ok {
			log.Errorln(ginresponse.MissingRequiredQuery(c, "destination"))
			return
		}
	case valueobject.PaginationSort.Number():
		number, ok = c.GetQuery("number")
		if !ok {
			log.Errorln(ginresponse.MissingRequiredQuery(c, "number"))
			return
		}
	default:
	}

	c.Set("load", load)
	c.Set("origin", origin)
	c.Set("destination", destination)
	c.Set("number", number)
	c.Set("sortBy", sortBy)

	c.Next()
}
