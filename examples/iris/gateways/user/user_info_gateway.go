package user

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
)

func (p *server) RegisterUserinfo(party iris.Party) {
	g := party.Party("/userinfo")

	g.Get("/get", p.getUserInfo)
}

func (p *server) getUserInfo(c iris.Context) {
	reversed, err := p.Reverser.Get().Reverse(c, "userinfo")
	if err != nil {
		c.StopWithError(http.StatusInternalServerError, err)
		return
	}

	c.WriteString(fmt.Sprintf("Hello, %s!\n", reversed))
}
