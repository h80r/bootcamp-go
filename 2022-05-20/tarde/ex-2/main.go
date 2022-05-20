package main

import (
	product "ex-2/pkg/product"
	store "ex-2/pkg/store"
	"fmt"
)

func main() {
	hStore := store.NovaLoja()
	hStore.Adicionar(product.NovoProduto(
		product.KTamanhoPequeno,
		"Google Pixel 6",
		5999.0,
	))

	hStore.Adicionar(product.NovoProduto(
		product.KTamanhoMedio,
		"Dyson Air AM09",
		2628.0,
	))

	hStore.Adicionar(product.NovoProduto(
		product.KTamanhoGrande,
		"Geladeira Samsung French Door RF27T5501SG",
		27699.99,
	))

	fmt.Printf("Custo total: R$ %.2f\n", hStore.Total())
}
