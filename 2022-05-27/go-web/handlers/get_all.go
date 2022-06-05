package handlers

import (
	"net/http"

	"go-web/utils"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	users, err := utils.ReadDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deserializar usuários",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &users)
}
