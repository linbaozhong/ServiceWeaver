package main

import (
	"context"
	"examples/iris/handlers"
	"github.com/ServiceWeaver/weaver"
	"github.com/kataras/iris/v12"
)

type Server struct {
	weaver.Implements[weaver.Main]
	hello weaver.Listener

	h weaver.Ref[handlers.IHello]
	u weaver.Ref[handlers.IUser]
}

func run(ctx context.Context, server *Server) error {
	//e := server.Init(ctx)
	//if e != nil {
	//	server.Logger(ctx).Error(e.Error())
	//}
	//return e
	return nil
}

func (p *Server) Init(ctx context.Context) error {
	app := iris.New().Configure(iris.WithRemoteAddrHeader(
		"X-Real-Ip",
		"X-Forwarded-For",
	))

	// 调试服务 Prepare for commissioning services
	app.Get("/", debug)
	app.Head("/", debug)

	// 注册路由 Registered route
	v1 := app.Party("/v1")

	appCtx := context.WithValue(ctx, "party", v1)

	p.h.Get().RegisterRouter(appCtx)
	p.u.Get().RegisterRouter(appCtx)

	e := app.Run(iris.Listener(p.hello),
		iris.WithLogLevel("debug"))
	if e != nil {
		p.Logger(ctx).Error(e.Error())
	}

	return nil

}

func debug(c iris.Context) {

	c.JSON(iris.Map{
		"name":    "WeaverServices(iris) Hello",
		"version": "0.1.0",
	})
}
