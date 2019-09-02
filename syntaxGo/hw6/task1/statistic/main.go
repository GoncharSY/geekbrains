package statistic

// Average - вернет среднее значение.
func Average(xs []float64) float64 {
	var total = float64(0)

	for i := range xs {
		total += xs[i]
	}

	return total / float64(len(xs))
}

// Sum - вернет сумму чисел в переданном наборе.
func Sum(itm []float64) (sum float64) {
	for i := range itm {
		sum += itm[i]
	}

	return
}
