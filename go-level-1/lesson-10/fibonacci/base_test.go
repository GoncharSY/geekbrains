package fibonacci

import (
	"testing"
)

func TestCalculateNumber(t *testing.T) {
	var cases = []struct {
		Input  int
		Output int
		IsErr  bool
		Descr  string
	}{
		{
			Input:  -1,
			Output: 0,
			IsErr:  true,
			Descr:  "test with input -1",
		}, {
			Input:  0,
			Output: 0,
			IsErr:  false,
			Descr:  "test with input 0",
		}, {
			Input:  1,
			Output: 1,
			IsErr:  false,
			Descr:  "test with input 1",
		}, {
			Input:  2,
			Output: 1,
			IsErr:  false,
			Descr:  "test with input 2",
		}, {
			Input:  3,
			Output: 2,
			IsErr:  false,
			Descr:  "test with input 3",
		}, {
			Input:  4,
			Output: 3,
			IsErr:  false,
			Descr:  "test with input 4",
		}, {
			Input:  5,
			Output: 5,
			IsErr:  false,
			Descr:  "test with input 5",
		},
	}

	for _, cs := range cases {
		val, err := CalculateNumber(cs.Input)

		if err != nil && !cs.IsErr {
			t.Error(
				"For", cs.Descr,
				"got error:", err.Error(),
			)
		} else if err == nil && cs.IsErr {
			t.Error(
				"For", cs.Descr,
				"expected error",
			)
		} else if val != cs.Output {
			t.Error(
				"For", cs.Descr,
				"expected", cs.Output,
				"got", val,
			)
		}
	}
}

func BenchmarkCalculateNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateNumber(20)
	}
}
