package user

import (
	"github.com/kataras/iris/v12"
)

type server struct {
}

var Server server

func (p *server) Register(party iris.Party, args ...any) {
	User.RegisterRouter(party, args...)
	Userinfo.RegisterRouter(party, args...)
}
