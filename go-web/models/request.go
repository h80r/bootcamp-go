package models

import "fmt"

type Request struct {
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Email   string  `json:"email"`
	Age     int     `json:"age"`
	Height  float64 `json:"height"`
}

func (r *Request) IsValid() error {
	// 2. Caso algum campo não esteja completo, um código de erro 400 deve ser retornado
	// com a mensagem “campo %s é obrigatório”.
	// (Em %s deve ir o nome do campo que não está completo).
	if r.Name == "" {
		return fmt.Errorf("campo name é obrigatório")
	}

	if r.Surname == "" {
		return fmt.Errorf("campo surname é obrigatório")
	}

	if r.Email == "" {
		return fmt.Errorf("campo email é obrigatório")
	}

	if r.Age == 0 {
		return fmt.Errorf("campo age é obrigatório")
	}

	if r.Height == 0 {
		return fmt.Errorf("campo height é obrigatório")
	}

	return nil
}
