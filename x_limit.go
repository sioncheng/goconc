package main

import (
	"context"
	"fmt"

	"golang.org/x/time/rate"
)

/**
APIConnection with rate limiter
*/
type APIConnection struct {
	rateLimiter *rate.Limiter
}

/**
Open open APIConnection
*/
func Open() *APIConnection {
	return &APIConnection{
		rateLimiter: rate.NewLimiter(rate.Limit(1), 1),
	}
}

/**
Read read
*/
func (a *APIConnection) Read(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	fmt.Println("Read...")

	return nil
}

func main() {

	a := Open()
	c, cf := context.WithCancel(context.TODO())

	for i := 0; i < 5; i++ {
		a.Read(c)
		if i == 2 {
			cf()
		}
	}
}
