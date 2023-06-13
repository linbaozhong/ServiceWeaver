package routers

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/gin-gonic/gin"
	"hello/components/reverse"
	"hello/handlers"
	"net/http"
)

type T interface {
	InitRouter(ctx context.Context) error
}
type router struct {
	weaver.Implements[T]
	reverser weaver.Ref[reverse.T]
}

func (r *router) InitRouter(ctx context.Context) error {
	app := gin.Default()

	// 调试服务 Prepare for commissioning services
	app.GET("/", debug)
	app.HEAD("/", debug)

	g := app.Group("/")
	// 注册路由 Registered route
	l := len(handlers.Instances)
	for i := 0; i < l; i++ {
		if m, ok := handlers.Instances[i].(handlers.IRegisterRouter); ok {
			m.RegisterRouter(g, r.reverser.Get())
		}
	}

	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, e := r.Listener("hello", opts)
	if e != nil {
		r.Logger().Error(e.Error())
		return e
	}

	e = app.RunListener(lis)
	if e != nil {
		r.Logger().Error(e.Error())
	}

	return nil
}

func debug(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"name":    "WeaverServices(iris) Hello",
		"version": "0.1.0",
	})
}
