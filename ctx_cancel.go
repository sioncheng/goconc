package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	fmt.Println(ctx)

	ctx2, cancelFunc := context.WithCancel(ctx)

	fmt.Println(ctx2, cancelFunc)

	go func(ctx context.Context) {
	loop:
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel")
				break loop
			case <-time.After(1 * time.Second):
				fmt.Println("work")

			}
		}
	}(ctx2)

	time.Sleep(10 * time.Second)

	cancelFunc()
}
