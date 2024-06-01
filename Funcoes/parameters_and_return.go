package main

import (
	"fmt"
)

func parameters(name string) {
	fmt.Println("Hello, my name is", name)
}

func returnNumber(x int) int {
	return x * x
}

func main() {
	parameters("Jackson")
	result := returnNumber(10)
	fmt.Println("The result is", result)
}
