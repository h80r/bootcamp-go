package main

import "fmt"

// Em sua função “main”, defina uma variável chamada “salario” e atribua um valor do
// tipo “int”.
func main() {
	salario := 1500
	msg, err := ValidateSalary(salario)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(msg)
	}
}

func ValidateSalary(salary int) (string, error) {
	if salary < 15000 {
		return "", &SalaryError{}
	}

	// Caso contrário, imprima no
	// console a mensagem “Necessário pagamento de imposto”.
	return "Necessário pagamento de imposto", nil
}

// 2. Crie um erro personalizado com uma struct que implemente “Error()” com a
// mensagem “erro: O salário digitado não está dentro do valor mínimo" em que seja
// disparado quando “salário” for menor que 15.000.
type SalaryError struct{}

func (e *SalaryError) Error() string {
	return "erro: O salário digitado não está dentro do valor mínimo"
}
