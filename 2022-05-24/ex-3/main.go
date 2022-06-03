package main

import "fmt"

// Precisamos de 3 estruturas:
// - Produtos: nome, preço, quantidade.
type Product struct {
	Name   string
	Price  float64
	Amount int
}

// - Serviços: nome, preço, minutos trabalhados.
type Service struct {
	Name    string
	Price   float64
	Minutes int
}

// - Manutenção: nome, preço.
type Maintenance struct {
	Name  string
	Price float64
}

// Precisamos de 3 funções:
// - Somar Produtos: recebe um array de produto e devolve o preço total (preço *
// quantidade).
func SumProducts(products []Product, out chan<- float64) {
	total := 0.0
	for _, product := range products {
		fmt.Println("Somando produto:", product.Name)
		total += product.Price * float64(product.Amount)
	}

	out <- total
}

// - Somar Serviços: recebe uma array de serviço e devolve o preço total (preço * média
// hora trabalhada, se ele não vier trabalhar por 30 minutos, ele será cobrado como se
// tivesse trabalhado meia hora).
func SumServices(services []Service, out chan<- float64) {
	total := 0.0
	for _, service := range services {
		fmt.Println("Somando serviço:", service.Name)
		total += service.Price * float64(func(minutes int) int {
			if minutes < 30 {
				return 1
			} else {
				return minutes / 30
			}
		}(service.Minutes))
	}

	out <- total
}

// - Somar Manutenção: recebe um array de manutenção e devolve o preço total.
func SumMaintenance(maintenances []Maintenance, out chan<- float64) {
	total := 0.0
	for _, maintenance := range maintenances {
		fmt.Println("Somando manutenção:", maintenance.Name)
		total += maintenance.Price
	}

	out <- total
}

func main() {
	products := []Product{
		{Name: "Coca-Cola", Price: 1.5, Amount: 2},
		{Name: "Sprite", Price: 2.0, Amount: 1},
		{Name: "Fanta", Price: 1.5, Amount: 1},
	}

	services := []Service{
		{Name: "Corte", Price: 10.0, Minutes: 27},
		{Name: "Limpeza", Price: 20.0, Minutes: 60},
		{Name: "Manutenção", Price: 50.0, Minutes: 120},
	}

	maintenances := []Maintenance{
		{Name: "Carro", Price: 100.0},
		{Name: "Caminhão", Price: 200.0},
		{Name: "Moto", Price: 50.0},
	}

	results := make(chan float64)
	go SumProducts(products, results)
	go SumServices(services, results)
	go SumMaintenance(maintenances, results)

	sum := <-results
	sum += <-results
	sum += <-results

	fmt.Println("Total:", sum)
}
