package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := media(7, 8, 9, -10, 11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Média: %.2f\n", res)
	}
}

func media(notas ...int) (float64, error) {
	total := 0

	for _, nota := range notas {
		if nota < 0 {
			return 0, errors.New("nota inválida")
		}
		total += nota
	}

	return float64(total) / float64(len(notas)), nil
}
