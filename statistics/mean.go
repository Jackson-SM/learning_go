package statistics

func Mean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}

	result := sum / float64(len(data))

	return result
}