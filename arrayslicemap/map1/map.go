package main

import "fmt"

func main() {

	aprovados := make(map[int]string)

	aprovados[12345678] = "Maria"
	aprovados[12345648] = "Pedro"
	aprovados[12345628] = "Ana"
	aprovados[12345618] = "João"

	for cpf, nome := range aprovados {

		fmt.Printf("%s (CPF: %d)\n  ", nome, cpf)
	}

	fmt.Println(aprovados[12345678])

	delete(aprovados, 12345618)

	fmt.Println(aprovados)
}
