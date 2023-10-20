package middleware

import (
	"errors"
	"fmt"
	"github.com/achillescres/pkg/gin/ginresponse"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"strings"
)

func (m *middleware) Jwt(c *gin.Context) error {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		log.Errorln(errEmptyHeader)
		ginresponse.ErrorString(c, http.StatusUnauthorized, errEmptyHeader, "empty auth header")
		return errEmptyHeader
	}

	headerParts := strings.Fields(strings.TrimSpace(header))
	if len(headerParts) != 2 || headerParts[0] != bearer {
		log.Errorln(errInvalidHeader)
		ginresponse.ErrorString(c, http.StatusUnauthorized, errInvalidHeader, "invalid auth header")
		return errInvalidHeader
	}

	if headerParts[1] == "" {
		log.Errorln(errInvalidHeader)
		ginresponse.ErrorString(c, http.StatusUnauthorized, errInvalidHeader, "token is empty")
		return errInvalidHeader
	}
	token := headerParts[1]

	userLogin, err := m.policyService.CheckUserToken(c, token)
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusUnauthorized, err, "token data is wrong")
		return err
	}

	user, err := m.policyService.GetUserByLogin(c, userLogin)
	if !user.Approved {
		err := errors.New("error unapproved user")
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "unapproved user")
		return err
	}
	if err != nil {
		log.Errorln(err)
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "couldn't get user")
		return err
	}
	if user == nil {
		ginresponse.ErrorString(c, http.StatusInternalServerError, err, "there's no such user")
		return err
	}

	c.Set(userKey, *user)
	c.Set(userLoginKey, userLogin)

	return nil
}

func GetUserLogin(c *gin.Context) (string, error) {
	_login, ok := c.Get(userLoginKey)
	if !ok {
		return "", fmt.Errorf("user login not found")
	}
	login, ok := _login.(string)
	if !ok {
		return "", fmt.Errorf("couldn't assert user login to string")
	}
	return login, nil
}

func GetUser(c *gin.Context) (*entity.User, error) {
	_user, ok := c.Get(userKey)
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	user, ok := _user.(entity.User)
	if !ok {
		return nil, fmt.Errorf("couldn't assert user to its type")
	}

	return &user, nil
}
