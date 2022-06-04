package main

import (
	"fmt"
	"os"
)

// Um escritório de contabilidade precisa acessar os dados de seus funcionários para
// poder realizar diferentes liquidações. Para isso, eles têm todos os detalhes
// necessários em um arquivo .txt.
// 1. É necessário desenvolver a funcionalidade para poder ler o arquivo .txt
// indicado pelo cliente, porém, eles não passaram o arquivo para ser lido pelo
// nosso programa.
func main() {
	// 2. Desenvolva o código necessário para ler os dados do arquivo chamado
	// “customers.txt” (lembre-se do que você viu sobre o pacote “os”).
	bytes, err := os.ReadFile("customers.txt")

	// Como não temos o arquivo necessário, será obtido um erro e, neste caso, o programa
	// terá que exibir um panic ao tentar ler um arquivo que não existe, exibindo a
	// mensagem “o arquivo indicado não foi encontrado ou está danificado ”.
	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}

	fmt.Println(string(bytes))
}
