package handlers

import "github.com/kataras/iris/v12"

type IRegisterRouter interface {
	RegisterRouter(party iris.Party, t ...interface{})
}

var (
	Instances = make([]interface{}, 0)
)
