package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Name     string
	Surname  string
	Email    string
	Products []Product
}

type Product struct {
	Name   string
	Price  float64
	Amount int
}

func NewProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price}
}

func AddProduct(user *User, product *Product, amount int) {
	product.Amount = amount
	user.Products = append(user.Products, *product)
}

func DeleteProducts(user *User) {
	user.Products = nil
}

func (user *User) String() string {
	products := (*productList)(&user.Products)
	return "Name: " + user.Name + "\nSurname: " + user.Surname + "\nEmail: " + user.Email + "\nProducts: " + products.String()
}

type productList []Product

func (products *productList) String() string {
	var str string
	for _, product := range *products {
		str += product.String()
	}
	return str
}

func (product *Product) String() string {
	return product.Name + " - " + strconv.FormatFloat(product.Price, 'f', 2, 64) + " - " + strconv.Itoa(product.Amount) + "\n"
}

func main() {
	user := User{
		Name:    "João",
		Surname: "Silva",
		Email:   "joao@gmail.com",
	}

	nintendoSwitch := NewProduct("Nintendo Switch", 2990)
	fmt.Println(nintendoSwitch)
	AddProduct(&user, nintendoSwitch, 1)
	fmt.Println(nintendoSwitch)

	tabS8 := NewProduct("Samsung Galaxy Tab S8", 2999)
	AddProduct(&user, tabS8, 2)

	fmt.Println(user)

	DeleteProducts(&user)

	fmt.Println(user)
}
