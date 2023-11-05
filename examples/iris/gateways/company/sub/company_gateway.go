package sub

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"iris/components/reverse"
	"net/http"
)

type company struct {
	reverser reverse.Reverser
}

var (
	Company company
)

func (p *company) RegisterRouter(party iris.Party, args ...any) {
	for _, arg := range args {
		switch a := arg.(type) {
		case reverse.Reverser:
			p.reverser = a
		}
	}
	g := party.Party("/company")
	g.Get("/", Company.hello)
	g.Get("/hi", Company.hi)
}

func (p *company) hello(c iris.Context) {
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

func (p *company) hi(c iris.Context) {
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
