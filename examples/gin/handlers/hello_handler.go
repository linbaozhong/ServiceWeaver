package handlers

import (
	"context"
	"github.com/ServiceWeaver/weaver/examples/gin/components/reverse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type hello struct {
	reverser reverse.T
}

func init() {
	Instances = append(Instances, &hello{})
}

func (p *hello) RegisterRouter(party *gin.RouterGroup, ts ...interface{}) {
	if len(ts) == 1 {
		p.reverser = ts[0].(reverse.T)
	}

	g := party.Group("/hello")
	g.GET("/", p.hello)
	g.GET("/hi", p.hi)
}

func (p *hello) hello(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "World"
	}
	reversed, err := p.reverser.Reverse(context.Background(), name)
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
	reversed, err := p.reverser.Reverse(context.Background(), name)
	if err != nil {
		c.String(http.StatusInternalServerError, "[ERROR]:", err)
		return
	}

	c.String(http.StatusOK, "Hi, %s!\n", reversed)

}
