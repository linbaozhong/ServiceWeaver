package handlers

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IUser interface {
	RegisterRouter(ctx context.Context) error
}

type user struct {
	weaver.Implements[IUser]
}

func (p *user) RegisterRouter(ctx context.Context) error {
	party := ctx.Value("party").(*gin.RouterGroup)

	g := party.Group("/user")
	g.GET("/", p.get)

	return nil
}

func (p *user) get(c *gin.Context) {
	c.String(http.StatusOK, "user OK !!!")
}
