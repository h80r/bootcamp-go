package handlers

import (
	"encoding/json"
	"os"

	"go-web/models"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
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

	var parsedFile map[string]interface{}
	err = json.Unmarshal(fileBytes, &parsedFile)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Erro ao deserializar arquivo",
			"error":   err.Error(),
		})
		return
	}

	// 4. Crie um slice da estrutura e retorne-a por meio de nosso endpoint.
	var users []models.User
	for _, user := range parsedFile["data"].([]interface{}) {
		var newUser models.User

		err = mapstructure.Decode(user, &newUser)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Erro ao decodificar usuário",
				"error":   err.Error(),
			})
			return
		}

		users = append(users, newUser)
	}

	c.JSON(200, users)
}
