package user

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"iris/components/reverse"
	"net/http"
)

type userinfo struct {
	reverser reverse.Reverser
}

var Userinfo userinfo

func (p *userinfo) RegisterRouter(party iris.Party, args ...any) {
	for _, arg := range args {
		switch a := arg.(type) {
		case reverse.Reverser:
			p.reverser = a
		}
	}
	g := party.Party("/userinfo")

	g.Get("/", p.get)
}

func (p *userinfo) get(c iris.Context) {
	reversed, err := p.reverser.Reverse(c, "userinfo")
	if err != nil {
		c.StopWithError(http.StatusInternalServerError, err)
		return
	}

	c.WriteString(fmt.Sprintf("Hello, %s!\n", reversed))
}
