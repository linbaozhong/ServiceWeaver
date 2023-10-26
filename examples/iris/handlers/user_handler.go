package handlers

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/kataras/iris/v12"
)

type IUser interface {
	RegisterRouter(ctx context.Context) error
}

type user struct {
	weaver.Implements[IUser]
}

func (p *user) RegisterRouter(ctx context.Context) error {
	party := ctx.Value("party").(iris.Party)

	g := party.Party("/user")
	g.Get("/", p.get)

	return nil
}

func (p *user) get(c iris.Context) {
	c.WriteString("user OK !!!")
}
