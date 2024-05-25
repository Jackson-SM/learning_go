package main

import "fmt"

func main() {
	// the ":=" operator is known like the short declaration operator or "gopher" operator
	// this operator do the declaration and the assignment of the variable, but "=" only do the assignment
	x := 10
	y := "Hello, Good morning!"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T", z, z)
}
