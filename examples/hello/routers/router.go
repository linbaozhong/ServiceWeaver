package routers

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/kataras/iris/v12"
	"hello/blls/reverse"
	"hello/handlers"
)

type T interface {
	InitRouter(ctx context.Context) error
}
type router struct {
	weaver.Implements[T]
	reverser weaver.Ref[reverse.T]
}

func (r *router) InitRouter(ctx context.Context) error {
	app := iris.New().Configure(iris.WithRemoteAddrHeader(
		"X-Real-Ip",
		"X-Forwarded-For",
	))

	// 调试服务
	app.Get("/", debug)
	app.Head("/", debug)
	//

	//注册路由
	handlers.Hello(r.reverser.Get(), app)

	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, e := r.Listener("hello", opts)
	if e != nil {
		return e
	}
	////
	//idleConnsClosed := make(chan struct{})
	//iris.RegisterOnInterrupt(func() {
	//	timeout := 10 * time.Second
	//	c, cancel := context.WithTimeout(ctx, timeout)
	//	defer cancel()
	//	// close all hosts.
	//	app.Shutdown(c)
	//	close(idleConnsClosed)
	//})

	e = app.Run(iris.Listener(lis),
		iris.WithLogLevel("debug"),
		iris.WithoutInterruptHandler)

	//<-idleConnsClosed
	return e
}

func debug(c iris.Context) {

	c.JSON(iris.Map{
		"name":    "weaver services",
		"version": "0.1.0",
	})

	return
}
