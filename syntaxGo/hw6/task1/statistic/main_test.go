package statistic

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverageSet(t *testing.T) {
	for i := range tests {
		v := Average(tests[i].values)
		if v != tests[i].average {
			t.Error(
				"For", tests[i].values,
				"expected", tests[i].average,
				"got", v,
			)
		}
	}
}
