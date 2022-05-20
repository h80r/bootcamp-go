package main

import (
	"errors"
	"fmt"
)

const (
	kCachorro  = "cachorro"
	kGato      = "gato"
	kHamster   = "hamster"
	kTarantula = "tarantula"
)

const (
	kCachorroCome  = 10000
	kGatoCome      = 5000
	kHamsterCome   = 250
	kTarantulaCome = 150
)

func main() {
	animalCachorro, erro := Animal(kCachorro)
	if erro != nil {
		fmt.Println(erro)
	}

	animalGato, erro := Animal(kGato)
	if erro != nil {
		fmt.Println(erro)
	}

	totalComida := animalCachorro(2) + animalGato(3)

	fmt.Printf("Total de comida: %.2fg\n", totalComida)
}

func calcularGastoCachorro(cachorros int) float64 {
	return float64(cachorros) * kCachorroCome
}

func calcularGastoGato(gatos int) float64 {
	return float64(gatos) * kGatoCome
}

func calcularGastoHamster(hamsters int) float64 {
	return float64(hamsters) * kHamsterCome
}

func calcularGastoTarantula(tarantulas int) float64 {
	return float64(tarantulas) * kTarantulaCome
}

func Animal(animal string) (func(int) float64, error) {
	switch animal {
	case kCachorro:
		return calcularGastoCachorro, nil
	case kGato:
		return calcularGastoGato, nil
	case kHamster:
		return calcularGastoHamster, nil
	case kTarantula:
		return calcularGastoTarantula, nil
	}

	return nil, errors.New("animal inválido")
}
