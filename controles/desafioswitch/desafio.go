package main

import "fmt"

func notaParaConceito(n float64) string {

	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"

	}
}

func main() {

	fmt.Println(notaParaConceito(9))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(4))
	fmt.Println(notaParaConceito(1))
}
