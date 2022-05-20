package product

type produto struct {
	tipo  int
	nome  string
	preco float64
}

type Produto interface {
	CalcularCusto() float64
}

type ProdutoPequeno struct {
	p produto
}

func (p ProdutoPequeno) CalcularCusto() float64 {
	return p.p.preco*(1+kTaxaEstoquePequeno) + kCustoEnvioPequeno
}

type ProdutoMedio struct {
	p produto
}

func (m ProdutoMedio) CalcularCusto() float64 {
	return m.p.preco*(1+kTaxaEstoqueMedio) + kCustoEnvioMedio
}

type ProdutoGrande struct {
	p produto
}

func (g ProdutoGrande) CalcularCusto() float64 {
	return g.p.preco*(1+kTaxaEstoqueGrande) + kCustoEnvioGrande
}

func NovoProduto(tipo int, nome string, preco float64) Produto {
	switch tipo {
	case KTamanhoPequeno:
		return ProdutoPequeno{p: produto{tipo: tipo, nome: nome, preco: preco}}
	case KTamanhoMedio:
		return ProdutoMedio{p: produto{tipo: tipo, nome: nome, preco: preco}}
	case KTamanhoGrande:
		return ProdutoGrande{p: produto{tipo: tipo, nome: nome, preco: preco}}
	default:
		return nil
	}
}
