package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/dangelzm/CinemaNetworkApi/controllers/TodoController"
	"github.com/dangelzm/CinemaNetworkApi/controllers/UserController"
)

type Controller interface {
	New(routerGroup *gin.RouterGroup)
}

var list = map[string]Controller{
	"UserController": new(UserController.Controller),
	"TodoController": new(TodoController.Controller),
}

func Mount(routerGroup *gin.RouterGroup, controller string) {
	list[controller].New(routerGroup)
}
