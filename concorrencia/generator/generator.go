package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

//Google IO/2012 - GO GOcurrency Patterns

func titulo(urls ...string) <-chan string {

	c := make(chan string)

	for _, url := range urls {

		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func main() {
	t1 := titulo("https://www.uol.com.br", "https://www.terra.com.br")
	t2 := titulo("https://g1.globo.com", "https://www.r7.com")

	fmt.Println("Primeiros: ", <-t1, "|", <-t2)
	fmt.Println("Segundos: ", <-t1, "|", <-t2)
}
