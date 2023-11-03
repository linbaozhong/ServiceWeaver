package user

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
)

func (p *server) RegisterUser(party iris.Party) {
	g := party.Party("/user")

	g.Get("/hello", p.hello)
	g.Get("/hi", p.hi)
}

func (p *server) hello(c iris.Context) {
	name := c.FormValue("name")
	if name == "" {
		name = "World"
	}
	reversed, err := p.reverser.Get().Reverse(context.Background(), name)
	if err != nil {
		c.StopWithError(http.StatusInternalServerError, err)
		return
	}

	c.WriteString(fmt.Sprintf("Hello, %s!\n", reversed))
}

func (p *server) hi(c iris.Context) {
	name := c.FormValue("name")
	if name == "" {
		name = "World"
	}

	reversed, err := p.reverser.Get().Reverse(context.Background(), name)
	if err != nil {
		c.StopWithError(http.StatusInternalServerError, err)
		return
	}

	c.WriteString(fmt.Sprintf("Hi, %s!\n", reversed))
}
