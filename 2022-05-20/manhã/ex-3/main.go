package main

import "fmt"

const (
	kCategoriaC = iota
	kCategoriaB
	kCategoriaA
)

const (
	kSalarioCategoriaC = 1000.0
	kSalarioCategoriaB = 1500.0
	kSalarioCategoriaA = 3000.0
)

const (
	kBonusCategoriaC = 0.0
	kBonusCategoriaB = 0.2
	kBonusCategoriaA = 0.5
)

const kHorasParaBonus = 160

func main() {
	var funcionario1 = map[string]int{
		"categoria":        kCategoriaC,
		"horasTrabalhadas": 162,
	}

	funcionario2 := map[string]int{
		"categoria":        kCategoriaB,
		"horasTrabalhadas": 176,
	}

	funcionario3 := map[string]int{
		"categoria":        kCategoriaA,
		"horasTrabalhadas": 172,
	}

	fmt.Printf("Funcionário 1: %.2f\n", salario(funcionario1["categoria"], funcionario1["horasTrabalhadas"]))
	fmt.Printf("Funcionário 2: %.2f\n", salario(funcionario2["categoria"], funcionario2["horasTrabalhadas"]))
	fmt.Printf("Funcionário 3: %.2f\n", salario(funcionario3["categoria"], funcionario3["horasTrabalhadas"]))
}

func salario(categoria int, horasTrabalhadas int) float64 {
	var salario float64 = 0.0

	switch categoria {
	case kCategoriaC:
		salario = kSalarioCategoriaC * float64(horasTrabalhadas)
	case kCategoriaB:
		salario = kSalarioCategoriaB * float64(horasTrabalhadas)
	case kCategoriaA:
		salario = kSalarioCategoriaA * float64(horasTrabalhadas)
	}

	if horasTrabalhadas > kHorasParaBonus {
		switch categoria {
		case kCategoriaB:
			salario += kBonusCategoriaB * salario

		case kCategoriaA:
			salario += kBonusCategoriaA * salario
		}
	}

	return salario
}
