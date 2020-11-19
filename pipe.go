package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()

	return out
}

func main() {
	for n := range gen(2, 3, 4) {
		fmt.Println(n)
	}

	for m := range sq(gen(1, 2, 3, 4, 5)) {
		fmt.Println(m)
	}
}
