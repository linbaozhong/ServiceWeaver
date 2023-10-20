package handlers

import (
	"github.com/kataras/iris/v12"
)

const (
	V1 = "v1"
	V2 = "v2"
)

type IRegisterRouter interface {
	RegisterRouter(party iris.Party)
}

var (
	Instances = make([]interface{}, 0)
)
