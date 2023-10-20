package main

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/examples/iris/routers"
)

type Server struct {
	weaver.Implements[weaver.Main]
	app weaver.Ref[routers.T]
}

func server(ctx context.Context, server *Server) error {
	e := server.app.Get().InitRouter(ctx)
	if e != nil {
		server.Logger().Error(e.Error())
	}
	return e
}
