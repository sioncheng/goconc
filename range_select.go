package main

import "fmt"

func main() {
	ch := make(chan interface{})

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
