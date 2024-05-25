package main

import "fmt"

func main() {
	bytenumbers, errors := fmt.Println("Writing a random text to return bytes and errors", "Hello World", 100) // this function returns the number of bytes and errors
	fmt.Println(bytenumbers, errors)                                                                           // printing the number of bytes and errors

	_, my_errors := fmt.Println("Writing a random text to return only errors", "Hello World", 100) // if we won't need to use the first return value, we can use _ to ignore it
	fmt.Println(my_errors)                                                                         // printing the errors

}
