package gateways

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/kataras/iris/v12"
	"iris/components/reverse"
)

type Server struct {
	weaver.Implements[weaver.Main]
	hello    weaver.Listener
	reverser weaver.Ref[reverse.Reverser]
}

func Run(ctx context.Context, server *Server) error {
	if e := server.start(ctx); e != nil {
		server.Logger(ctx).Error(e.Error())
		return e
	}
	return nil
}

func (p *Server) start(ctx context.Context) error {
	Hello.Register()

	e := App().Run(iris.Listener(p.hello),
		iris.WithLogLevel("debug"))
	if e != nil {
		p.Logger(ctx).Error(e.Error())
	}

	return nil
}
