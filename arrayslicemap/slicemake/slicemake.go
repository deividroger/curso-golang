package main

import "fmt"

func main() {

	s := make([]int, 10)
	s[9] = 12

	fmt.Println(s)

	s = make([]int, 10, 20) //array interno tem 20 posições, e 10 itens alocados

	fmt.Println(s, len(s), cap(s)) //caps retorna a capacidade do array

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 2)

	fmt.Println(s, len(s), cap(s))
}
