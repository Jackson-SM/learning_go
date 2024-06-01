package main

import (
	"fmt"
	"runtime"
)

func main() {

	core := runtime.NumCPU()

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
		fmt.Println(core)
	}
}
