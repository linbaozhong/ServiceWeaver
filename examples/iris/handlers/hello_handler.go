package handlers

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/examples/iris/components/reverse"
	"github.com/kataras/iris/v12"
	"net/http"
)

type Hello interface {
	RegisterRouter(context.Context) error
}

type hi struct {
	weaver.Implements[Hello]
	reverser weaver.Ref[reverse.Reverser]
}

//func init() {
//	Instances = append(Instances, &hi{})
//}

func (p *hi) RegisterRouter(ctx context.Context) error {
	party, _ := ctx.Value(V1).(iris.Party)

	g := party.Party("/hello")
	g.Get("/", p.hello)
	g.Get("/hi", p.hi)

	return nil
}

func (p *hi) hello(c iris.Context) {
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

func (p *hi) hi(c iris.Context) {
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
