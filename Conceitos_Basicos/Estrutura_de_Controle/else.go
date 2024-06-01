package main

import "fmt"

func verifyAgeWithElse(age int) {
	if age <= 17 {
		fmt.Println("You are younger than 18 years old")
	} else {
		fmt.Println("You are older than 18 years old")
	}
}

func main() {
	verifyAgeWithElse(17)
	verifyAgeWithElse(18)
	verifyAgeWithElse(19)
}
