package main

import (
	"errors"
	"fmt"
)

const (
	kOperacaoMinimo = iota
	kOperacaoMedia
	kOperacaoMaximo
)

func main() {
	obterMinimo, erro := definirOperacao(kOperacaoMinimo)
	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Printf("Mínimo: %.2f\n", obterMinimo(7, 8, 9, 10, 11))
	}
}

func definirOperacao(operacao int) (func(...float64) float64, error) {
	switch operacao {
	case kOperacaoMinimo:
		return minimo, nil
	case kOperacaoMedia:
		return media, nil
	case kOperacaoMaximo:
		return maximo, nil
	}

	return nil, errors.New("operação inválida")
}

func minimo(notas ...float64) float64 {
	var min = notas[0]

	for _, nota := range notas {
		if nota < min {
			min = nota
		}
	}

	return min
}

func media(notas ...float64) float64 {
	var total = 0.0

	for _, nota := range notas {
		total += nota
	}

	return float64(total) / float64(len(notas))
}

func maximo(notas ...float64) float64 {
	var max = notas[0]

	for _, nota := range notas {
		if nota > max {
			max = nota
		}
	}

	return max
}
