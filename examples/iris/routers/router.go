package routers

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

type IRegisterRouter interface {
	RegisterRouter(party iris.Party, args ...any)
}

type app struct {
	name        string
	application *iris.Application
}

func NewApp(name string) *app {
	a := &app{
		name: name,
		application: iris.New().Configure(iris.WithRemoteAddrHeader(
			"X-Real-Ip",
			"X-Forwarded-For",
		)),
	}

	// 调试服务
	a.application.Get("/", a.debug)
	a.application.Head("/", a.debug)

	return a
}

func (p *app) debug(c iris.Context) {
	c.JSON(iris.Map{
		"name":    fmt.Sprintf("WeaverServices(iris) %s", p.name),
		"version": "0.1.0",
	})
}

func (p *app) Application() *iris.Application {
	return p.application
}

func (p *app) Name() string {
	return p.name
}
