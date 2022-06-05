package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("Please set token environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == requiredToken {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"code": http.StatusUnauthorized,
				"error": func() string {
					if token == "" {
						return "Token vazio"
					}
					return "Token inválido"
				}(),
			},
		)
	}
}
