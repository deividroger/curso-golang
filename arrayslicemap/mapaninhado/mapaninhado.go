package main

import "fmt"

func main() {

	funcsPorLetra := map[string]map[string]float64{

		"G": {
			"Gabriela Silva": 15251,
			"Guga Perera":    1522,
		},
		"P": {
			"Pedro": 1135,
		},
		"A": {
			"Aline F": 6510,
		},
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)

		for nome, valor := range funcs {

			fmt.Println(nome, valor)
		}

	}

}
