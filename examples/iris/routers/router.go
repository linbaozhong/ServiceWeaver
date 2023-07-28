package routers

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/examples/iris/components/reverse"
	"github.com/ServiceWeaver/weaver/examples/iris/handlers"
	"github.com/kataras/iris/v12"
)

type T interface {
	InitRouter(ctx context.Context) error
}
type router struct {
	weaver.Implements[T]
	reverser weaver.Ref[reverse.T]
	lis      weaver.Listener
}

func (r *router) InitRouter(ctx context.Context) error {
	fmt.Printf("hello listener available on %v\n", r.lis)
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

	e := app.Run(iris.Listener(r.lis),
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
