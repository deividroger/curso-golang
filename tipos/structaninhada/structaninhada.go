package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) valorTotal() float64 {

	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}

	return total
}

func main() {

	pedido := pedido{
		userID: 1,
		itens: []item{
			item{1, 2, 12.10}, //evite esse tipo de declaração
			item{
				produtoID: 2,
				qtde:      1,
				preco:     23.49,
			},
			item{11, 100, 3.13}, //evite esse tipo de declaração
		},
	}

	fmt.Printf("Valor Total do pedido é R$ %.2f  ", pedido.valorTotal())
}
