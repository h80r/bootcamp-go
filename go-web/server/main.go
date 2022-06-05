package main

import (
	"go-web/internal"
	"go-web/server/handlers"
	"go-web/server/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	repo := internal.NewRepository()
	service := internal.NewService(repo)
	controller := handlers.NewUserController(service)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá Heitor",
		})
	})

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controller.GetAll)
		userRoutes.GET("/:id", controller.GetById)

		userRoutes.Use(middleware.AuthMiddleware())

		userRoutes.POST("/", controller.Create)

		userRoutes.PUT("/:id", controller.Modify)
	}

	router.Run(":8080")
}
