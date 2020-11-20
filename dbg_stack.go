package main

import (
	"fmt"
	"runtime/debug"
)

func a() int {
	a := 1
	b := 2

	return a / b
}

func main() {
	c := a()
	fmt.Println(c)
	fmt.Println(string(debug.Stack()))
}
