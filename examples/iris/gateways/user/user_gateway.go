package user

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"iris/components/reverse"
	"net/http"
)

type user struct {
	reverser reverse.Reverser
}

var User user

func (p *user) Register(args ...interface{}) {
	for _, arg := range args {
		switch i := arg.(type) {
		case reverse.Reverser:
			p.reverser = i
		}
	}

	app.Get("/v1/"+name, User.hello)
	app.Get("/v1/"+name+"/hi", User.hi)
}

func (p *user) hello(c iris.Context) {
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

func (p *user) hi(c iris.Context) {
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
