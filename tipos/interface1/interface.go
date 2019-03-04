package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {

	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {

	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}

	fmt.Println(coisa.toString())

	coisa = produto{"Notebook", 1630.25}
	fmt.Println(coisa.toString())
}
