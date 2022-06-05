package internal

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"
)

type Repository interface {
	GetAll() ([]User, error)
	GetById(id int) (User, error)
	LastID() (int, error)
	Create(name, surname, email string, age int, height float64) (User, error)
}

type repository struct {
	database *[]User
}

func (r *repository) GetAll() ([]User, error) {
	if r.database != nil {
		return *r.database, nil
	}

	fileBytes, err := os.ReadFile("../users.json")
	if err != nil {
		return nil, err
	}

	var users []User
	if err := json.Unmarshal(fileBytes, &users); err != nil {
		return nil, err
	}

	r.database = &users

	return *r.database, nil
}

func (r *repository) GetById(id int) (User, error) {
	db, err := r.GetAll()
	if err != nil {
		return User{}, err
	}

	for _, user := range db {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, errors.New("users array does not contain user with id " + strconv.Itoa(id))
}

func (r *repository) LastID() (int, error) {
	db, err := r.GetAll()
	if err != nil {
		return 0, err
	}

	return db[len(db)-1].ID, nil
}

func (r *repository) Create(name, surname, email string, age int, height float64) (User, error) {
	newID, err := r.LastID()
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:        newID + 1,
		Name:      name,
		Surname:   surname,
		Email:     email,
		Age:       age,
		Height:    height,
		Active:    true,
		CreatedAt: time.Now().Format("2006-01-02"),
	}

	*r.database = append(*r.database, user)
	return user, nil
}

func NewRepository() Repository {
	return &repository{}
}
