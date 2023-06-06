package handlers

import (
	"context"
	"examples/hello/components"
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
)

type hello struct {
	//component weaver.Instance
	reverser components.Reverser
}

func Hello(reverser components.Reverser, party iris.Party) *hello {
	obj := &hello{
		reverser: reverser,
	}
	g := party.Party("/hello")
	g.Get("/", obj.hello)
	g.Get("/hi", obj.hi)

	return obj
}

func (p *hello) hello(c iris.Context) {
	name := c.FormValue("name")
	if name == "" {
		name = "World"
	}
	reversed, err := p.reverser.Reverse(context.Background(), name)
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
	reversed, err := p.reverser.Reverse(context.Background(), name)
	if err != nil {
		c.StopWithError(http.StatusInternalServerError, err)
		return
	}

	c.WriteString(fmt.Sprintf("HI, %s!\n", reversed))

}
