package statistic

import (
	"testing"
)

type testCase struct {
	values []float64
	result float64
}

var avgSet = []testCase{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for i := range avgSet {
		v := Average(avgSet[i].values)

		if v != avgSet[i].result {
			t.Error(
				"For", avgSet[i].values,
				"expected", avgSet[i].result,
				"got", v,
			)
		}
	}
}

var sumSet = []testCase{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
}

func TestSum(t *testing.T) {
	for i := range sumSet {
		res := Sum(sumSet[i].values)

		if res != sumSet[i].result {
			t.Error(
				"For", sumSet[i].values,
				"expected", sumSet[i].result,
				"got", res,
			)
		}
	}
}
