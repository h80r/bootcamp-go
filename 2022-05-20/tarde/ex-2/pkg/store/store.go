package store

import (
	p "ex-2/pkg/product"
)

type Ecommerce interface {
	Total() float64
	Adicionar(p.Produto)
}

type loja struct {
	produtos []p.Produto
}

func (l loja) Total() float64 {
	total := 0.0
	for _, p := range l.produtos {
		total += p.CalcularCusto()
	}
	return total
}

func (l *loja) Adicionar(produto p.Produto) {
	l.produtos = append(l.produtos, produto)
}

func NovaLoja() Ecommerce {
	return &loja{}
}
