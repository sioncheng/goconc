package main

import "fmt"

func main() {
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{}, 1)

	select {
	case a := <-c1:
		fmt.Println("a", a)
	case b := <-c2:
		fmt.Println("b", b)
	case c3 <- 1:
		fmt.Println("c", 1)
	}
}
