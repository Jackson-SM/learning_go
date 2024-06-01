package main

import "fmt"

func main() {
	name := "John"
	var age int = 20

	var street, city string = "123 Main St", "San Francisco"
	var country, zip = "USA", 94122

	const s string = "constant"

	fmt.Println("Hello, my name is", name, "and I am", age, "years old.")
	fmt.Println("I am a", s, "string.")
	fmt.Println("I live in", city, "on", street)
	fmt.Println("The zip code is", zip, "and the country is", country)
}
