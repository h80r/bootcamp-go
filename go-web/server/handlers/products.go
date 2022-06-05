package handlers

import (
	"go-web/internal"
	"go-web/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service internal.Service
}

func (u *UserController) Create(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "0908762111" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Você não tem permissão para fazer a solicitação solicitada",
		})
		return
	}

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

	newUser, err := u.service.Create(req.Name, req.Surname, req.Email, req.Age, req.Height)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Não foi possível criar o usuário.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso.",
		"user":    newUser,
	})
}

func (u *UserController) GetAll(c *gin.Context) {
	users, err := u.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Não foi possível obter os usuários.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (u *UserController) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	user, err := u.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Usuário não encontrado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// É solicitado a implementação de uma funcionalidade que modifique completamente uma
// entidade. Para isso, é necessário seguir os seguintes passos:
// 1. Gere um método PUT para modificar toda a entidade.
// 2. No Path envie o ID da entidade a ser modificada.
// 3. Se não existir, retorne um erro 404..
// 4. Realize todas as validações (todos os campos são obrigatórios).
func (u *UserController) Modify(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "0908762111" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Você não tem permissão para fazer a solicitação solicitada",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	_, err = u.service.UserExists(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Usuário não encontrado",
			"error":   err.Error(),
		})
		return
	}

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

	user, err := u.service.Modify(id, req.Name, req.Surname, req.Email, req.Age, req.Height)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Não foi possível modificar o usuário.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func NewUserController(service internal.Service) *UserController {
	return &UserController{
		service: service,
	}
}
