package TodoController

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/dangelzm/cinema-network-api/models"
)

type Controller struct{}

var todos = map[string]models.TodoModel{
	"1": models.TodoModel{Id: 1, Title: "Test1"},
	"2": models.TodoModel{Id: 2, Title: "Test2"},
	"3": models.TodoModel{Id: 3, Title: "Test3"},
}

func (ctrl *Controller) New(r *gin.RouterGroup) {
	r.GET("/", getList)
	r.GET("/:id", getById)
}

func getList(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func getById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, todos[id])
}
