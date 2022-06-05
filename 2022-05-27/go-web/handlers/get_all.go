package handlers

import (
	"encoding/json"
	"os"

	"go-web/models"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	fileBytes, err := os.ReadFile("users.json")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Erro ao ler arquivo",
			"error":   err.Error(),
		})
		return
	}

	var users []models.User
	err = json.Unmarshal(fileBytes, &users)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Erro ao deserializar arquivo",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, users)
}
