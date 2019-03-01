package main

import "fmt"

func inc1(n int) {

}

func inc2(n *int) {

	*n++
}

func main() {
	n := 1

	fmt.Println(n)

	inc1(n)

	inc2(&n)

	fmt.Println(n)
}
