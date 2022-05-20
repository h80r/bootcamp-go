package main

import "fmt"

func main() {
	var salario = 149999.0

	fmt.Printf("Salário: %.2f\n", salario)
	fmt.Printf("Imposto: %.2f\n", impostoAplicado(salario))
}

func impostoAplicado(salario float64) float64 {
	if salario >= 150000 {
		return salario * 0.27
	}

	return salario * 0.17
}
