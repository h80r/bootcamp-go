package main

import "fmt"

func main() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Manuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	fmt.Printf("Idade de Benjamin: %d anos\n", employees["Benjamin"])

	qntdMaiorIdade := 0
	for _, idade := range employees {
		if idade > 21 {
			qntdMaiorIdade++
		}
	}
	fmt.Printf("Quantidade de pessoas com mais de 21 anos: %d\n", qntdMaiorIdade)

	employees["Federico"] = 25
	fmt.Println("Federico adicionado ao mapa:")
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println("Pedro removido do mapa:")
	fmt.Println(employees)
}
