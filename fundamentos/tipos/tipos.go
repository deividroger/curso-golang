package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//Números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	//sem sinal (só possitivos)

	var b byte = 255

	fmt.Println("O byte é", reflect.TypeOf(b))

	i1 := math.MaxInt64

	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo do i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela unicode (int32)

	fmt.Println("O rune é", reflect.TypeOf(i2))

	fmt.Println(i2)

	//números reais (float32, float64)

	var x = float32(49.99)

	fmt.Println("O tipo de x é", reflect.TypeOf(x))

	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	//boolean
	bo := true

	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))

	fmt.Println(!bo)

	s1 := "Olá meu nome é Deivid"

	fmt.Println("O tipo de bo é", reflect.TypeOf(s1))
	fmt.Println("O tamanho da string é", len(s1))
	fmt.Println(s1)

	char := 'a'

	fmt.Println("O tipo de char é", reflect.TypeOf(char))

	fmt.Println(char)
}
