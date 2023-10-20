package routers

import (
	"context"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/examples/gin/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type T interface {
	InitRouter(ctx context.Context) error
}

type router struct {
	weaver.Implements[T]
	lis weaver.Listener
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
			m.RegisterRouter(g)
		}
	}

	e := app.RunListener(r.lis)
	if e != nil {
		r.Logger(ctx).Error(e.Error())
	}

	return nil
}

func debug(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"name":    "WeaverServices(gin) Hello",
		"version": "0.1.0",
	})
}
