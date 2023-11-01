package gateways

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
)

type hello struct {
	Server
}

var Hello hello

func (p *hello) Register() {
	App().Get("/v1/hello", Hello.hello)
	App().Get("/v1/hello/hi", Hello.hi)
}

func (p *hello) hello(c iris.Context) {
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

func (p *hello) hi(c iris.Context) {
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
