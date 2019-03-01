package main

import "fmt"

func main() {

	//array pré-fixado
	numeros := [...]int{1, 2, 3, 4, 5}

	//aqui eu uso o indice e o valor
	for i, numero := range numeros {

		fmt.Printf("%d %d\n", i, numero)
	}

	// aqui eu uso apenas o valor
	for _, num := range numeros {
		fmt.Println(num)
	}

}
