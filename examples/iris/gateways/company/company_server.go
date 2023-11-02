package company

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/kataras/iris/v12"
	"iris/components/reverse"
	"iris/gateways/company/sub"
	"iris/routers"
)

var (
	App  *iris.Application
	name string = "company"
)

type Server interface {
	Run(ctx context.Context) error
	Shutdown(ctx context.Context) error
}

type server struct {
	weaver.Implements[Server]
	company weaver.Listener

	Reverser weaver.Ref[reverse.Reverser]
}

func (p *server) Init(ctx context.Context) error {
	a := routers.NewApp(name)
	App = a.Application()

	return nil
}

func (p *server) Run(ctx context.Context) error {
	v1 := App.Party("/v1")

	sub.User.Register(v1, p.Reverser.Get())

	e := App.Run(iris.Listener(p.company),
		iris.WithLogLevel("debug"))
	if e != nil {
		p.Logger(ctx).Error(e.Error())
	}

	return nil
}

func (p *server) Shutdown(ctx context.Context) error {
	fmt.Println("已关闭：", name)
	return App.Shutdown(ctx)
}

func (p *server) GetReverser() reverse.Reverser {
	return p.Reverser.Get()
}
