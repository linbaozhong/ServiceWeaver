package handlers

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"hello/components/reverse"
	"net/http"
)

type hello struct {
	reverser reverse.T
}

func init() {
	Instances = append(Instances, &hello{})
}

func (p *hello) RegisterRouter(party iris.Party, ts ...interface{}) {
	if len(ts) == 1 {
		p.reverser = ts[0].(reverse.T)
	}

	g := party.Party("/hello")
	g.Get("/", p.hello)
	g.Get("/hi", p.hi)
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

	c.WriteString(fmt.Sprintf("Hi, %s!\n", reversed))
}
