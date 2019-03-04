package main

import (
	"fmt"

	"github.com/deividroger/html"
)

func encaminhar(origem <-chan string, destino chan string) {

	for {
		destino <- <-origem
	}
}

//multiplexar - misturar (mensagens) no mesmo canal

func juntar(entrada1, entrada2 <-chan string) <-chan string {

	c := make(chan string)

	go encaminhar(entrada1, c)

	go encaminhar(entrada2, c)

	return c
}

func main() {

	c := juntar(
		html.Titulo("https://www.uol.com.br", "https://www.terra.com.br"),
		html.Titulo("https://g1.globo.com", "https://www.r7.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
