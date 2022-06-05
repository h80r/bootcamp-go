package main

import (
	"go-web/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá Heitor",
		})
	})

	router.GET("/users", handlers.GetAll)
	router.GET("/users/:id", handlers.GetById)

	// A funcionalidade para criar a entidade deve ser implementada. Se isso acontecer, os
	// seguintes passos devem ser seguidos:
	// 1. Crie um endpoint por meio de POST que receba a entidade.
	router.POST("/users", handlers.Create)

	router.Run(":8080")
}
