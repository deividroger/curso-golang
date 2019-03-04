package main

import "fmt"

func main() {

	ch := make(chan int, 1)

	//como eu mando e recebo dados com channels?

	ch <- 1 //enviando dads para o channel (escrita)

	<-ch //recebendo dados do channl (leitura)

	ch <- 2

	fmt.Println(<-ch)
}
