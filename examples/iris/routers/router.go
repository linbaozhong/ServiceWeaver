package routers

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/kataras/iris/v12"
	"hello/components/reverse"
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

	// 调试服务 Prepare for commissioning services
	app.Get("/", debug)
	app.Head("/", debug)

	// 注册路由 Registered route
	l := len(handlers.Instances)
	for i := 0; i < l; i++ {
		if m, ok := handlers.Instances[i].(handlers.IRegisterRouter); ok {
			m.RegisterRouter(app, r.reverser.Get())
		}
	}

	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, e := r.Listener("hello", opts)
	if e != nil {
		r.Logger().Error(e.Error())
		return e
	}

	e = app.Run(iris.Listener(lis),
		iris.WithLogLevel("debug"))
	if e != nil {
		r.Logger().Error(e.Error())
	}

	return nil
}

func debug(c iris.Context) {

	c.JSON(iris.Map{
		"name":    "WeaverServices(iris) Hello",
		"version": "0.1.0",
	})
}
