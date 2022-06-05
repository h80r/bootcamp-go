package handlers

import (
	"go-web/internal"
	"go-web/models"
	"go-web/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	// Para adicionar segurança à aplicação, o pedido deve ser enviado com um token, para isso
	// devem ser seguidos os seguintes passos:
	// 1. No momento do envio da solicitação, deve ser validado que um token é enviado
	// 2. Esse token deve ser validado em nosso código (o token pode ser codificado
	// permanentemente).
	// 3. Caso o token enviado não esteja correto, devemos retornar um erro 401 e uma
	// mensagem que "você não tem permissão para fazer a solicitação solicitada".
	token := c.GetHeader("Authorization")
	if token != "0908762111" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Você não tem permissão para fazer a solicitação solicitada",
		})
		return
	}

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

	if err := req.IsValid(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Não foi possível validar a requisição.",
			"error":   err.Error(),
		})
		return
	}

	newUser := internal.User{
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
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso.",
		"user":    newUser,
	})
}
