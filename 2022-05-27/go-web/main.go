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

	router.Run(":8080")
}
