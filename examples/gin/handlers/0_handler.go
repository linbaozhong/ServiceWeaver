package handlers

import "github.com/gin-gonic/gin"

type IRegisterRouter interface {
	RegisterRouter(party *gin.RouterGroup)
}

var (
	Instances = make([]interface{}, 0)
)
