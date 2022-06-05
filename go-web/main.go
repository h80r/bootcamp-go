package main

import (
	"go-web/handlers"

	"github.com/gin-gonic/gin"
)

// A estrutura do projeto deve ser separada e como primeiro passo gerando o pacote interno,
// todas as funcionalidades que não dependem de pacotes externos devem estar no pacote
// interno.

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá Heitor",
		})
	})

	router.GET("/users", handlers.GetAll)
	router.GET("/users/:id", handlers.GetById)

	router.POST("/users", handlers.Create)

	router.Run(":8080")
}
