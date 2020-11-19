package main

import (
	"fmt"
	"sync"
)

func g(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

func merge(chans ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	fanin := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(chans))
	for _, c := range chans {
		go fanin(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	for n := range merge(g(1, 2), g(3, 4), g(5, 6)) {
		fmt.Println(n)
	}
}
