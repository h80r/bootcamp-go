package handlers

import (
	"go-web/models"
	"go-web/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 3. No momento de fazer a solicitação, o ID deve ser gerado. Para gerar o ID, devemos
// procurar o ID do último registro gerado, incrementá-lo em 1 e atribuí-lo ao nosso novo
// registro (sem ter uma variável global do último ID).
func Create(c *gin.Context) {
	database, err := utils.ReadDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Não foi possível ler o banco de dados.",
			"error":   err.Error(),
		})
		return
	}

	lastID := (*database)[len(*database)-1].ID

	var req models.Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Não foi possível ler o JSON.",
			"error":   err.Error(),
		})
		return
	}

	newUser := models.User{
		ID:        lastID + 1,
		Name:      req.Name,
		Surname:   req.Surname,
		Email:     req.Email,
		Age:       req.Age,
		Height:    req.Height,
		Active:    true,
		CreatedAt: time.Now().Format("2006-01-02"),
	}

	utils.IncludeDatabase(&newUser)
}
