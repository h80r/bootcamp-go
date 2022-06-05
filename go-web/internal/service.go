package internal

type Service interface {
	GetAll() ([]User, error)
	GetById(id int) (User, error)
	Create(name, surname, email string, age int, height float64) (User, error)
	UserExists(id int) error
	Modify(id int, name, surname, email string, age int, height float64) (User, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return *users, nil
}

func (s *service) GetById(id int) (User, error) {
	user, err := s.repository.GetById(id)
	if err != nil {
		return User{}, err
	}
	return *user, nil
}

func (s *service) Create(name, surname, email string, age int, height float64) (User, error) {
	user, err := s.repository.Create(name, surname, email, age, height)
	if err != nil {
		return User{}, err
	}
	return *user, nil
}

func (s *service) Modify(id int, name, surname, email string, age int, height float64) (User, error) {
	user, err := s.repository.Modify(id, name, surname, email, age, height)
	if err != nil {
		return User{}, err
	}
	return *user, nil
}

func (s *service) UserExists(id int) error {
	return s.repository.UserExists(id)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
