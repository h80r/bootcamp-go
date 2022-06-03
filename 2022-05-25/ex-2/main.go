package main

import (
	"errors"
	"fmt"
)

// Faça a mesma coisa que no exercício anterior, mas reformule o código para que, em vez de
// “Error()”, seja implementado “errors.New()”.

func main() {
	salario := 1500
	msg := ValidateSalary(salario)
	fmt.Println(msg.Error())
}

func ValidateSalary(salary int) error {
	if salary < 15000 {
		return errors.New("erro: O salário digitado não está dentro do valor mínimo")
	}

	return errors.New("necessário pagamento de imposto")
}
