package statistics

func Moda(data []float64) float64 {
	frequency := make(map[float64]int)
	max := 0
	var result float64

	for _, value := range data {
		frequency[value]++

		if frequency[value] > max {
			max = frequency[value]
			result = value
		}
	}

	return result
}