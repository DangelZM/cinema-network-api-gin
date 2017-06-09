package UserController

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {}

func (ctrl *Controller ) New(r *gin.RouterGroup) {
	r.GET("/", getList)
}

func getList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Users List Here",
	})
}
