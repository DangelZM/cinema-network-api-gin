package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct{}

func (this *UserController) New(r *gin.RouterGroup) {
	r.GET("/", this.getList)
	r.GET("/:id", this.getById)
}

func (ctrl *UserController) getList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Users List Here",
	})
}

func (ctrl *UserController) getById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "User with " + id + " here",
	})
}
