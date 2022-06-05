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

func (u *UserController) Modify(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	err = u.service.UserExists(id)
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

// É necessário implementar uma funcionalidade para excluir uma entidade. Para isso, é
// necessário seguir os seguintes passos:
// 1. Gere um método DELETE para excluir a entidade com base no ID.
// 2. Se não existir, retorne um erro 404.
func (u *UserController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	err = u.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Usuário não encontrado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário deletado com sucesso.",
	})
}

func NewUserController(service internal.Service) *UserController {
	return &UserController{
		service: service,
	}
}
