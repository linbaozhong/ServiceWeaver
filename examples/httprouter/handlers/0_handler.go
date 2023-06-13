package handlers

import (
	"github.com/julienschmidt/httprouter"
)

type IRegisterRouter interface {
	RegisterRouter(party *httprouter.Router, t ...interface{})
}

var (
	Instances = make([]interface{}, 0)
)
