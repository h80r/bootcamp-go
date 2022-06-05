package main

import (
	"go-web/internal"
	"go-web/server/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := internal.NewRepository()
	service := internal.NewService(repo)
	controller := handlers.NewUserController(service)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá Heitor",
		})
	})

	router.GET("/users", controller.GetAll)
	router.GET("/users/:id", controller.GetById)

	router.POST("/users", controller.Create)

	router.PUT("/users/:id", controller.Modify)

	router.Run(":8080")
}
