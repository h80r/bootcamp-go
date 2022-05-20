package main

import "fmt"

func main() {
	var idade = 21
	var estaEmpregado = true
	var anosTrabalhados = 0.0
	var salario = 7000.0

	fmt.Println(processarCliente(idade, estaEmpregado, float32(anosTrabalhados), salario))
}

func processarCliente(idade int, estaEmpregado bool, anosTrabalhados float32, salario float64) string {
	if idade <= 22 {
		return "[Inelegível] Você precisa ser mais velho para receber este empréstimo!"
	}

	if !estaEmpregado {
		return "[Inelegível] Você precisa estar empregado para receber este empréstimo!"
	}

	if anosTrabalhados <= 1.0 {
		return "[Inelegível] Você precisa trabalhar mais de 1 ano para receber este empréstimo!"
	}

	if salario > 100000 {
		return "[Elegível] Você pode receber o empréstimo sem juros!"
	}

	return "[Elegível] Você pode receber o empréstimo com juros!"
}
