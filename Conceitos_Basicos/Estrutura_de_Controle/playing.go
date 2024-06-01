package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if number <= 5 {
			fmt.Println("This number is less than or equal to 5")
			break
		}

		println("I'm inside the loop", i)
	}
}
