package handlers

import (
	"go-web/models"
	"go-web/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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
			"message": "Não foi possível ler a requisição.",
			"error":   err.Error(),
		})
		return
	}

	// As validações dos campos devem ser implementadas no momento do envio do pedido, para
	// isso devem ser seguidos os seguintes passos:
	// 1. Todos os campos enviados na solicitação devem ser validados, todos os campos são
	// obrigatórios
	if err := req.IsValid(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Não foi possível validar a requisição.",
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
