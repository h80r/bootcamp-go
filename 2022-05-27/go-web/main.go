// Já tendo criado e testado nossa API que nos recebe, geramos uma rota que retorna uma lista
// do tema escolhido.

package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// 1. Dentro do "main.go", crie uma estrutura de acordo com o tema com os campos
// correspondentes.
// "Usuários variam por id, nome, sobrenome, e-mail, idade, altura, ativo (sim-não), data de criação.",
type User struct {
	ID        int     `mapstructure:"id"`
	Name      string  `mapstructure:"name"`
	Surname   string  `mapstructure:"surname"`
	Email     string  `mapstructure:"email"`
	Age       int     `mapstructure:"age"`
	Height    float64 `mapstructure:"height"`
	Active    bool    `mapstructure:"active"`
	CreatedAt string  `mapstructure:"created_at"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá Heitor",
		})
	})

	// 2. Crie um endpoint cujo caminho é /users
	router.GET("/users", GetAll)

	router.Run(":8080")
}

// 3. Crie uma handler para o endpoint chamada "GetAll".
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
	var users []User
	for _, user := range parsedFile["data"].([]interface{}) {
		var newUser User

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
