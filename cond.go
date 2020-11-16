package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		c.L.Unlock()
		fmt.Println("remove from queue...")
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 1 {
			c.Wait()
		}
		queue = append(queue, struct{}{})
		fmt.Println("add to queue...")
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
