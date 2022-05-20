package main

import (
	"fmt"
)

func main() {
	var palavra = "Academia"
	fmt.Println(len(palavra))
	for _, letra := range palavra {
		fmt.Println(string(letra))
	}
}
