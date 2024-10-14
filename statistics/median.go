package statistics

import (
	"slices"
)

func Median(data []float64) float64 {
	slices.Sort(data)
	length := len(data)
	if length%2 == 0 {
		return (data[length/2-1] + data[length/2]) / 2
	}
	return data[length/2]
}