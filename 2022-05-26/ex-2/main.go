package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

// O mesmo escritório do exercício anterior, solicita uma funcionalidade para poder
// cadastrar dados de novos clientes. Os dados necessários para cadastrar um cliente são:
// - Arquivo
// - Nome e sobrenome
// - RG
// - Número de telefone
// - Endereço
type Register struct {
	FileName string
	Name     string
	Surname  string
	RG       string
	Phone    string
	Address  string
}

// ● Tarefa 1: O número do arquivo deve ser atribuído ou gerado separadamente e antes
// da cobrança das despesas restantes. Desenvolva e implemente uma função para
// gerar um ID que você usará posteriormente para atribuí-lo como um valor a “Arquivo”.
// Se por algum motivo esta função retornar o valor "nil", deve gerar um panic que
// interrompa a execução e aborte
func FileNameGenerator() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return id.String()
}

// Tarefa 2: Antes de cadastrar um cliente, você deve verificar se o cliente já existe. Para
// fazer isso, você precisa ler os dados de um arquivo .txt. Em algum lugar do seu
// código, implemente a função para ler um arquivo chamado “customers.txt” (como no
// exercício anterior, este arquivo não existe, então qualquer função que tente lê-lo
// retornará um erro). Você deve lidar adequadamente com esse erro, como vimos até agora.
// Esse erro deve:
// 1. Gerar um panic;
// 2. Imprimir no console a mensagem: “erro: o arquivo indicado não foi encontrado ou
// está danificado”, e continuar com a execução do programa normalmente.
// Requisitos gerais:
// ● Lembre-se de realizar as validações necessárias para cada retorno que possa conter
// um valor de erro (por exemplo, aqueles que tentam ler arquivos).
func IsClientRegistered(rg string) bool {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("erro: o arquivo indicado não foi encontrado ou está danificado")
		}
	}()

	bytes, err := os.ReadFile("customers.txt")
	if err != nil {
		panic(err)
	}

	lines := string(bytes)
	return strings.Contains(lines, rg)
}

// Requisitos gerais:
// Crie um erro, personalizando-o ao seu gosto, utilizando qualquer uma das funções
// que o GO disponibiliza para ele (ele também deve realizar a validação relevante para
// o caso de erro retornado).
type RegistrationError struct{}

func (r *RegistrationError) Error() string {
	return "erro: o cliente já está cadastrado"
}

// Requisitos gerais:
// ● Use recover para recuperar o valor dos panics que podem surgir (exceto na tarefa 1).
func main() {
	defer func() {
		err := recover()
		if err == nil {
			return
		}

		var errorType *RegistrationError
		if errors.As(err.(error), &errorType) {
			fmt.Println(err)
			return
		}

		panic(err)
	}()

	register := Register{
		FileName: FileNameGenerator(),
		Name:     "João",
		Surname:  "Silva",
		RG:       "123456789",
		Phone:    "123456789",
		Address:  "Rua dos Bobos, 0",
	}

	isRegistered := IsClientRegistered(register.RG)
	if isRegistered {
		panic(&RegistrationError{})
	}

	fmt.Println(register)
}
