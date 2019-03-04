package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

//método: função com reveiver :(receptor)

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	var produto1 produto

	produto1 = produto{
		nome:     "Lápis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1989.80, 0.10}

	fmt.Println(produto2.precoComDesconto())

}
