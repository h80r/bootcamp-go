package main

import (
	"fmt"
	"os"
)

type Product struct {
	Id     int
	Price  float64
	Amount int
}

func (p Product) String() string {
	return fmt.Sprintf("%d;%.2f;%d", p.Id, p.Price, p.Amount)
}

func main() {
	var products []Product
	products = append(products, Product{111223, 30012.0, 1})
	products = append(products, Product{444321, 1000000.0, 4})
	products = append(products, Product{434321, 50.0, 1})

	var fileData = "ID;Preco;Quantidade"

	var total float64
	for _, product := range products {
		total += product.Price * float64(product.Amount)
		fileData += "\n" + product.String()
	}

	fileData += "\n;" + fmt.Sprintf("%.2f", total) + ";"

	bytes := []byte(fileData)
	err := os.WriteFile("products.csv", bytes, 0644)
	if err != nil {
		panic(err)
	}
}
