package main

import "fmt"

/*
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.
*/

func add(x int, y int) int {
	return x + y
}

/*
When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
*/
func addOmitTypes(x, y int) int {
	return x + y
}

/*
A function can return any number of results.
*/

func returnInfos(x, y string) (string, string) {
	return y, x
}

func main() {
	result := add(10, 20)
	resultOmit := addOmitTypes(10, 20)

	fmt.Println("The result is", result)
	fmt.Println("The result is", resultOmit)

	first, second := returnInfos("Jackson", "Magalhães")

	fmt.Println("My name is", second, "and my last name is", first)
}
