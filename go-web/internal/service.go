// Dentro do pacote devem estar as camadas:
// 1. Service, deve conter a lógica da nossa aplicação.
// a. O arquivo service.go deve ser criado.
// b. A interface Service com todos os seus métodos deve ser gerada.
// c. A estrutura de serviço que contém o repositório deve ser gerada.
// d. Deve ser gerada uma função que retorne o Serviço.
// e. Todos os métodos correspondentes às operações a serem executadas (GetAll,
// Create, etc.) devem ser implementados.

package internal

type Service interface {
	GetAll() ([]User, error)
	GetById(id int) (User, error)
	Create(name, surname, email string, age int, height float64) (User, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]User, error) {
	return s.repository.GetAll()
}

func (s *service) GetById(id int) (User, error) {
	return s.repository.GetById(id)
}

func (s *service) Create(name, surname, email string, age int, height float64) (User, error) {
	return s.repository.Create(name, surname, email, age, height)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
