package main

import "fmt"

func main() {
	i := 1

	var p *int

	p = &i //pegando o endereço da variável

	*p++ //desreferenciando (pegando o valor)

	//Go não tem aritmética de ponteiros
	//p++

	fmt.Println(p, *p, i, &i)

}
