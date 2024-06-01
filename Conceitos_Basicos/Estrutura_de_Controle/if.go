package main

import "fmt"

func verifyAge(age int) {
	if age <= 17 {
		fmt.Println("You are younger than 18 years old")
		return
	}

	fmt.Println("You are older than 18 years old")
}

func main() {
	verifyAge(17)
	verifyAge(18)
	verifyAge(19)
}
