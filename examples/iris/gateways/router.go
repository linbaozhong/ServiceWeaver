package gateways

import (
	"github.com/kataras/iris/v12"
	"sync"
)

var (
	application *iris.Application
	once        sync.Once
)

func App() *iris.Application {
	once.Do(func() {
		application = iris.New().Configure(iris.WithRemoteAddrHeader(
			"X-Real-Ip",
			"X-Forwarded-For",
		))

		// 调试服务
		application.Get("/", debug)
		application.Head("/", debug)
	})
	return application
}

func debug(c iris.Context) {
	c.JSON(iris.Map{
		"name":    "WeaverServices(iris) Hello",
		"version": "0.1.0",
	})
}
