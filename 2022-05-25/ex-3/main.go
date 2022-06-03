package main

import "fmt"

// Repita o processo anterior, mas agora implementando "fmt.Errorf()", para que a mensagem de
// erro receba como parâmetro o valor de "salario", indicando que não atinge o mínimo
// tributável (a mensagem exibida pelo console deve dizer : "erro: o mínimo tributável é 15.000 e
// o salário informado é: [salario]”, onde [salario] é o valor do tipo int passado pelo parâmetro).

func main() {
	salario := 15000
	msg := ValidateSalary(salario)
	fmt.Println(msg.Error())

}

func ValidateSalary(salary int) error {
	if salary < 15000 {
		return fmt.Errorf("erro: o mínimo tributável é 15.000 e o salário informado é: %d", salary)
	}

	return fmt.Errorf("necessário pagamento de imposto")
}
