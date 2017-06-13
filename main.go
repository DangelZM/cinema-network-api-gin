package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"runtime"

	"fmt"
	"github.com/dangelzm/cinema-network-api/controllers"
	"github.com/dangelzm/cinema-network-api/db"
	"github.com/dangelzm/cinema-network-api/middlewares"
	"time"
)

const (
	Port = "8080"
)

func init() {
	db.Connect()
}

func main() {
	app := gin.Default()

	// Middlewares
	app.Use(middlewares.Connect)
	app.Use(middlewares.ErrorHandler)
	app.Use(middlewares.CORS)

	app.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"apiVersion": "0.0.1",
			"goVersion":  runtime.Version(),
		})
	})

	controllers.Mount(app.Group("api/users"), "UserController")
	controllers.Mount(app.Group("api/todos"), "TodoController")

	app.NoRoute()

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}

	fmt.Println("Start listening on " + port)

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
