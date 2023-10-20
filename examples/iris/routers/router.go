package routers

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/examples/iris/handlers"
	"github.com/kataras/iris/v12"
)

type T interface {
	InitRouter(ctx context.Context) error
}
type router struct {
	weaver.Implements[T]
	hi    weaver.Ref[handlers.Hello]
	user  weaver.Ref[handlers.User]
	hello weaver.Listener
}

func (r *router) InitRouter(ctx context.Context) error {
	fmt.Printf("service listener available on %v\n", r.hello)
	app := iris.New().Configure(iris.WithRemoteAddrHeader(
		"X-Real-Ip",
		"X-Forwarded-For",
	))

	// 调试服务 Prepare for commissioning services
	app.Get("/", debug)
	app.Head("/", debug)

	// 注册路由 Registered route
	v1 := context.WithValue(ctx, handlers.V1, app.Party("/v1"))
	r.hi.Get().RegisterRouter(v1)
	r.user.Get().RegisterRouter(v1)

	e := app.Run(iris.Listener(r.hello),
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
