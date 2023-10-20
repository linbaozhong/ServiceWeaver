package handlers

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/kataras/iris/v12"
)

type User interface {
	RegisterRouter(context.Context) error
}
type user struct {
	weaver.Implements[User]
}

func (p *user) RegisterRouter(ctx context.Context) error {
	party, _ := ctx.Value(V1).(iris.Party)

	g := party.Party("/user")
	g.Get("/", p.get)

	return nil
}

func (p *user) get(c iris.Context) {
	c.WriteString("user OK !!!")
}
