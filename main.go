package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"runtime"

	"github.com/dangelzm/cinema-network-api/controllers"
	"github.com/dangelzm/cinema-network-api/middlewares"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := gin.Default()
	app.Use(middlewares.CORS())

	app.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"apiVersion": "0.0.1",
			"goVersion":  runtime.Version(),
		})
	})

	controllers.Mount(app.Group("api/users"), "UserController")
	controllers.Mount(app.Group("api/todos"), "TodoController")

	app.NoRoute()
	app.Run(":" + port)
}
