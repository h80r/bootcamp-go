package main

import "fmt"

type Aluno struct {
	Nome           string
	Sobrenome      string
	RG             string
	DataDeAdmissao string
}

func (a Aluno) Log() {
	fmt.Printf("[Aluno(a) RG %s]\n", a.RG)
	fmt.Printf("\tNome: %s\n", a.Nome)
	fmt.Printf("\tSobrenome: %s\n", a.Sobrenome)
	fmt.Printf("\tData de Admissão: %s\n", a.DataDeAdmissao)
}

func main() {
	aluno := Aluno{
		Nome:           "João",
		Sobrenome:      "Silva",
		RG:             "123456789",
		DataDeAdmissao: "01/01/2018",
	}

	aluno.Log()

	fmt.Println()

	aluno2 := Aluno{
		Nome:           "Maria",
		Sobrenome:      "Santos",
		RG:             "987654321",
		DataDeAdmissao: "01/01/2018",
	}

	aluno2.Log()
}
