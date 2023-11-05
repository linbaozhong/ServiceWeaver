package company

import (
	"github.com/kataras/iris/v12"
	"iris/gateways/company/sub"
)

type server struct {
}

var Server server

func (p *server) Register(party iris.Party, args ...any) {
	sub.Company.RegisterRouter(party, args...)
}
