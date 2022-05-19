package main

import "fmt"

func main() {
	var temperatura = 27.0
	var humidade = 0.28
	var pressao = 1012

	fmt.Println(fmt.Sprintf("Temperatura (float): %f˚C", temperatura))
	fmt.Println(fmt.Sprintf("Humidade (float): %f%%", humidade * 100))
	fmt.Println(fmt.Sprintf("Pressão (int): %d hPa", pressao))
}