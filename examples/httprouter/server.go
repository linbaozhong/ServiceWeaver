package main

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/examples/httprouter/components/reverse"
	"github.com/ServiceWeaver/weaver/examples/httprouter/handlers"
	"github.com/julienschmidt/httprouter"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"net/http"
)

type Server struct {
	weaver.Implements[weaver.Main]
	reverser weaver.Ref[reverse.T]
	lis      weaver.Listener
}

func (server *Server) Main(ctx context.Context) error {
	fmt.Printf("hello listener available on %v\n", server.lis)

	e := server.InitRouter(ctx)
	if e != nil {
		server.Logger().Error(e.Error())
	}
	return e
}

func (server *Server) InitRouter(ctx context.Context) error {
	app := httprouter.New()

	// 调试服务 Prepare for commissioning services
	app.GET("/", debug)
	app.HEAD("/", debug)

	// 注册路由 Registered route
	l := len(handlers.Instances)
	for i := 0; i < l; i++ {
		if m, ok := handlers.Instances[i].(handlers.IRegisterRouter); ok {
			m.RegisterRouter(app, server.reverser.Get())
		}
	}
	otelHandler := otelhttp.NewHandler(app, "http")

	e := http.Serve(server.lis, otelHandler)
	if e != nil {
		server.Logger().Error(e.Error())
	}

	return nil
}

func debug(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprint(w, map[string]string{
		"name":    "WeaverServices(hrrprouter) Hello",
		"version": "0.1.0",
	})
}
