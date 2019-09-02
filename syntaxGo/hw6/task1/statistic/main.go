package statistic

// Average - вернет среднее значение.
func Average(xs []float64) float64 {
	var total = float64(0)

	for i := range xs {
		total += xs[i]
	}

	return total / float64(len(xs))
}
