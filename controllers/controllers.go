package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	New(routerGroup *gin.RouterGroup)
}

var list = map[string]Controller {
	"UserController": new(UserController),
	"TodoController": new(TodoController),
}

func Mount(routerGroup *gin.RouterGroup, controller string) {
	list[controller].New(routerGroup)
}
