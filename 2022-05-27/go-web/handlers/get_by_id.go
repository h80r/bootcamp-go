package handlers

import (
	"errors"
	"go-web/models"
	"go-web/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUser(users *[]models.User, id int) (*models.User, error) {
	for _, user := range *users {
		if user.ID == id {
			return &user, nil
		}
	}

	return nil, errors.New("users array does not contain user with id " + strconv.Itoa(id))
}

func GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID inválido",
			"error":   err.Error(),
		})
		return
	}

	users, err := utils.ReadDatabase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deserializar usuários",
			"error":   err.Error(),
		})
		return
	}

	user, err := getUser(users, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Usuário não encontrado",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
