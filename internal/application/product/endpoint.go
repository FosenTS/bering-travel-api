package product

import (
	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	controllers *Controllers
}

func NewEndpoint(cs *Controllers) *Endpoint {
	return &Endpoint{controllers: cs}
}

func (e *Endpoint) RegisterHandlersToGroup(r *gin.Engine) error {
	root := r.Group("/")
	e.controllers.RegisterHandlers(root)

	return nil
}
