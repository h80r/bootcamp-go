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
