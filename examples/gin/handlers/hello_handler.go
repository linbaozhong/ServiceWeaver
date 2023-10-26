package handlers

import (
	"context"
	"examples/gin/components/reverse"
	"github.com/ServiceWeaver/weaver"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IHello interface {
	RegisterRouter(ctx context.Context) error
}

type hello struct {
	weaver.Implements[IHello]
	reverser weaver.Ref[reverse.T]
}

func (p *hello) RegisterRouter(ctx context.Context) error {
	party := ctx.Value("party").(*gin.RouterGroup)

	g := party.Group("/hello")
	g.GET("/", p.hello)
	g.GET("/hi", p.hi)

	return nil
}

func (p *hello) hello(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "World"
	}
	reversed, err := p.reverser.Get().Reverse(context.Background(), name)
	if err != nil {
		c.String(http.StatusInternalServerError, "[ERROR]:", err)
		return
	}

	c.String(http.StatusOK, "Hello, %s!\n", reversed)
}

func (p *hello) hi(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "World"
	}
	reversed, err := p.reverser.Get().Reverse(context.Background(), name)
	if err != nil {
		c.String(http.StatusInternalServerError, "[ERROR]:", err)
		return
	}

	c.String(http.StatusOK, "Hi, %s!\n", reversed)

}
