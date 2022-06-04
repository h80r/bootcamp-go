package main

import (
	"errors"
	"fmt"
)

type TaxesError struct {
	salary float64
}

func (t *TaxesError) Error() string {
	return fmt.Sprintf("info: o mínimo tributável é 20,000.00 e o salário informado é: %.2f", t.salary)
}

func CalculateTaxes(totalSalary float64) (float64, error) {
	// - No caso de o salário mensal ser igual ou superior a R$ 20.000, 10% devem ser
	// descontados como imposto.
	if totalSalary >= 20000 {
		return totalSalary * 0.1, nil
	}
	return 0, &TaxesError{totalSalary}
}

// Vamos fazer com que nosso programa seja um pouco mais complexo e útil.
// 1. Desenvolva as funções necessárias para permitir que a empresa calcule:
// a) Salário mensal de um funcionário segundo a quantidade de horas trabalhadas.
// - A função receberá as horas trabalhadas no mês e o valor da hora como
// parâmetro.
// - Esta função deve retornar mais de um valor (salário calculado e erro).
func CalculateSalary(hoursWorked int, valuePerHour float64) (float64, error) {
	// Se o número de horas mensais inseridas for menor que 80 ou um número
	// negativo, a função deverá retornar um erro. Deve indicar "erro: o trabalhador
	// não pode ter trabalhado menos de 80 horas por mês".
	if hoursWorked < 80 {
		return 0, errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}

	totalSalary := float64(hoursWorked) * valuePerHour
	taxes, err := CalculateTaxes(totalSalary)
	return totalSalary - taxes, fmt.Errorf("erro: %w", err)
}

// 2. Desenvolva o código necessário para cumprir as funcionalidades solicitadas, usando
// “errors.New()”, “fmt.Errorf()” e “errors.Unwrap()”. Não esqueça de realizar as validações dos
// retornos de erro em sua função “main()”.
func main() {
	salary, err := CalculateSalary(80, 10)
	if err != nil {
		if errors.Unwrap(err) == nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(errors.Unwrap(err))
		}
	}
	fmt.Printf("Salário calculado (já descontado imposto): %.2f\n", salary)
}
