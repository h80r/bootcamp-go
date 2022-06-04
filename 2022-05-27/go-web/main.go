// 1. Crie dentro da pasta go-web um arquivo chamado main.go
package main

// 2. Crie um servidor web com Gin que retorne um JSON que tenha uma chave
// “mensagem” e diga Olá seguido do seu nome.
import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Olá Heitor",
		})
	})

	router.Run(":8080")
}

// 3. Acesse o end-point para verificar se a resposta está correta.
//   Acessado utilizando Insomnia: http://localhost:8080/
//   Exibiu o JSON: {"message":"Olá Heitor"}
