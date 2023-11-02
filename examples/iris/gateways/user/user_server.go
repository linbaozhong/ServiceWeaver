package user

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/kataras/iris/v12"
	"iris/components/reverse"
	"iris/routers"
)

var (
	app  *iris.Application
	name string = "user"
)

type Server interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type server struct {
	weaver.Implements[Server]
	user weaver.Listener

	Reverser weaver.Ref[reverse.Reverser]
}

func (p *server) Init(ctx context.Context) error {
	a := routers.NewApp(name)
	app = a.Application()

	return nil
}

func (p *server) Run(ctx context.Context) error {
	v1 := app.Party("/v1")
	p.RegisterUserinfo(v1)
	p.RegisterUser(v1)

	e := app.Run(iris.Listener(p.user),
		iris.WithLogLevel("debug"))
	if e != nil {
		p.Logger(ctx).Error(e.Error())
	}

	return nil
}
func (p *server) Shutdown(ctx context.Context) error {
	fmt.Println("已关闭：", name)
	return app.Shutdown(ctx)
}
