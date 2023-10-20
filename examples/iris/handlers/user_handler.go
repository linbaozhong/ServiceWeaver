package handlers

import (
	"github.com/kataras/iris/v12"
)

type user struct {
}

func init() {
	Instances = append(Instances, &user{})
}

func (p *user) RegisterRouter(party iris.Party) error {
	g := party.Party("/user")
	g.Get("/", p.get)

	return nil
}

func (p *user) get(c iris.Context) {
	c.WriteString("user OK !!!")
}
