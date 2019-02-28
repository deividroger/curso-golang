package main

import (
	"fmt"
	"math"
)

func main() {

	a := 4
	b := 2

	fmt.Println("Soma =>", a+b)

	fmt.Println("Subtração =>", a-b)

	fmt.Println("Divisão =>", a/b)

	fmt.Println("Multiplicação =>", a*b)

	fmt.Println("Módulo =>", a%b)

	//bitwise

	fmt.Println("E =>", a&b) //11 & 10 = 10

	fmt.Println("OU =>", a|b) //11 | 10 = 10

	fmt.Println("Xor =>", a^b) //11 ^ 10 = 10

	c := 3.0
	d := 2.0

	fmt.Println("Maior=>", math.Max(c, d))

	fmt.Println("Exponenciação=>", math.Pow(c, d))

}
