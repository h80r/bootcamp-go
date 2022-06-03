package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Name     string
	Surname  string
	Age      int
	Email    string
	Password string
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age int) {
	u.Age = age
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) String() string {
	return "User: " + u.Name + " " + u.Surname + " " + strconv.Itoa(u.Age) + " " + u.Email + " " + u.Password
}

func main() {
	user := User{}
	user.SetName("John")
	user.SetAge(30)
	user.SetEmail("john@gmail.com")
	user.SetPassword("12345")
	fmt.Println(user.String())
}
