package main

import (
	"context"
	"examples/gin/handlers"
	"github.com/ServiceWeaver/weaver"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	weaver.Implements[weaver.Main]
	hello weaver.Listener

	h weaver.Ref[handlers.IHello]
	u weaver.Ref[handlers.IUser]
}

func run(ctx context.Context, server *Server) error {
	e := server.Init(ctx)
	if e != nil {
		server.Logger(ctx).Error(e.Error())
	}
	return e
}

func (p *Server) Init(ctx context.Context) error {
	app := gin.Default()

	// 调试服务 Prepare for commissioning services
	app.GET("/", debug)
	app.HEAD("/", debug)

	// 注册路由 Registered route
	v1 := app.Group("/v1")

	appCtx := context.WithValue(ctx, "party", v1)

	p.h.Get().RegisterRouter(appCtx)
	p.u.Get().RegisterRouter(appCtx)

	//l := len(handlers.Instances)
	//for i := 0; i < l; i++ {
	//	if m, ok := handlers.Instances[i].(handlers.IRegisterRouter); ok {
	//		m.RegisterRouter(appCtx)
	//	} else {
	//		fmt.Println("hello:", m)
	//	}
	//}

	e := app.RunListener(p.hello)
	if e != nil {
		p.Logger(ctx).Error(e.Error())
	}

	return nil

}

func debug(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "WeaverServices(gin) Hello",
		"version": "0.1.0",
	})
}
