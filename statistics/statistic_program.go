package statistics

import (
	"fmt"
)

func Statistic_Program() {
	var volume int
	var data []float64

	fmt.Print("Enter the number of elements: ")
	fmt.Scan(&volume)

	for i := 0; i < volume; i++ {
		var element float64
		fmt.Printf("Enter the %d element: ", i+1)
		fmt.Scan(&element)
		data = append(data, element)
	}

	mean := Mean(data)
	fmt.Printf("Mean: %f\n", mean)
	moda := Moda(data)
	fmt.Printf("Moda: %f\n", moda)
	median := Median(data)
	fmt.Printf("Median: %f\n", median)
}