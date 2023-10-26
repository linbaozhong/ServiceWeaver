package handlers

import (
	"context"
	"examples/iris/components/reverse"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/kataras/iris/v12"
	"net/http"
)

type IHello interface {
	RegisterRouter(ctx context.Context) error
}

type hello struct {
	weaver.Implements[IHello]
	reverser weaver.Ref[reverse.Reverser]
}

func (p *hello) RegisterRouter(ctx context.Context) error {
	party, ok := ctx.Value("party").(iris.Party)
	if !ok {
		return nil
	}

	g := party.Party("/hello")
	g.Get("/", p.hello)
	g.Get("/hi", p.hi)

	return nil
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
