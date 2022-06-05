package main

import (
	"go-web/internal"
	"go-web/server/handlers"

	"github.com/gin-gonic/gin"
)

// A estrutura do projeto deve ser separada, como segundo passo deve ser gerado o pacote do
// servidor onde serão adicionadas as funcionalidades do projeto que dependem de pacotes
// externos e o principal do programa.

// Dentro do pacote deve estar:
// 1. O main do programa.
// a. O repositório, serviço e handler devem ser importados e injetados
// b. O roteador deve ser implementado para os diferentes endpoints
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

	router.Run(":8080")
}
