package handlers

import (
	"github.com/julienschmidt/httprouter"
)

type IRegisterRouter interface {
	RegisterRouter(party *httprouter.Router)
}

var (
	Instances = make([]interface{}, 0)
)
