package main

import "fmt"

func main() {
	meses := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
	var mesEscolhido = 1
	fmt.Println(mesEscolhido, "de", meses[mesEscolhido-1])
}
