package controllers

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"net/http"

	"github.com/dangelzm/cinema-network-api/models"
)

type UserController struct{}

func (this *UserController) New(r *gin.RouterGroup) {
	r.POST("/", this.add)
	r.GET("/", this.getList)
	r.GET("/:id", this.getById)
}

func (ctrl *UserController) add(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionUser).Insert(user)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func (ctrl *UserController) getList(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	users := []models.User{}
	err := db.C(models.CollectionUser).Find(nil).All(&users)
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) getById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "User with " + id + " here",
	})
}
