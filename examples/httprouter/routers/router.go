package routers

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/julienschmidt/httprouter"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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
	app := httprouter.New()

	// 调试服务 Prepare for commissioning services
	app.GET("/", debug)
	app.HEAD("/", debug)

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

	otelHandler := otelhttp.NewHandler(app, "http")

	e = http.Serve(lis, otelHandler)
	if e != nil {
		r.Logger().Error(e.Error())
	}

	return nil
}

func debug(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, map[string]string{
		"name":    "WeaverServices(iris) Hello",
		"version": "0.1.0",
	})
}
