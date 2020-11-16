package main

import (
	"fmt"
	"sync"
)

func main() {

	newInstance := func() interface{} {
		fmt.Println("creating new instance.")
		return struct{}{}
	}

	myPool := &sync.Pool{
		New: newInstance,
	}

	s := myPool.Get()
	fmt.Println(s)
	s2 := myPool.Get()
	fmt.Println(s2)

	myPool.Put(s)
	myPool.Put(s2)

	s3 := myPool.Get()
	fmt.Println(s3)
}
