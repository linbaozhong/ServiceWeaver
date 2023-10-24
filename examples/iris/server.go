package main

import (
	"context"
	"examples/gin/routers"
	"github.com/ServiceWeaver/weaver"
)

type Server struct {
	weaver.Implements[weaver.Main]
	app weaver.Ref[routers.T]
}

func server(ctx context.Context, server *Server) error {
	e := server.app.Get().InitRouter(ctx)
	if e != nil {
		server.Logger(ctx).Error(e.Error())
	}
	return e
}
