package middleware

import (
	"errors"
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/storage/dto"
)

const FlightIdentityKey = "flightIdentityKey"

func (m *middleware) FlightIdentityFromQueries(c *gin.Context) {
	airlCode, ok := c.GetQuery("airlCode")
	if !ok {
		log.Errorln(ginresponse.MissingRequiredQuery(c, "airlCode"))
		return
	}
	number, ok := c.GetQuery("number")
	if !ok {
		log.Errorln(ginresponse.MissingRequiredQuery(c, "number"))
		return
	}
	originIATA, ok := c.GetQuery("originIATA")
	if !ok {
		log.Errorln(ginresponse.MissingRequiredQuery(c, "originIATA"))
		return
	}
	destinationIATA, ok := c.GetQuery("destinationIATA")
	if !ok {
		log.Errorln(ginresponse.MissingRequiredQuery(c, "destinationIATA"))
		return
	}
	fltDate, ok := c.GetQuery("fltDate")
	if !ok {
		log.Errorln(ginresponse.MissingRequiredQuery(c, "fltDate"))
		return
	}

	fIdentity := dto.NewFlightIdentity(airlCode, number, originIATA, destinationIATA, fltDate)
	c.Set(FlightIdentityKey, *fIdentity)
	c.Next()
}

func GetFlightIdentity(c *gin.Context) (*dto.FlightIdentity, error) {
	fiAny, ok := c.Get(FlightIdentityKey)
	if !ok {
		return nil, errors.New("error no FlightIdentity found")
	}
	fi, ok := fiAny.(dto.FlightIdentity)
	if !ok {
		return nil, errors.New("error couldn't assert FlightIdentity")
	}

	return &fi, nil
}
