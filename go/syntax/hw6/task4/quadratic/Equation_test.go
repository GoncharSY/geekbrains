package quadratic

import "testing"

type testCaseEquation struct {
	Equ Equation
	Dsc float64
	Str string
	Rts []float64
}

var cases = []testCaseEquation{
	// Обычный случай
	{
		Equ: Equation{A: 1, B: -8, C: 12},
		Dsc: 16,
		Str: "x*x - 8*x + 12 = 0",
		Rts: []float64{2, 6},
	},
	// Случай с дробными корнями
	{
		Equ: Equation{A: 8, B: 2, C: -1},
		Dsc: 36,
		Str: "8*x*x + 2*x - 1 = 0",
		Rts: []float64{-0.5, 0.25},
	},
	// Случай с одним корем
	{
		Equ: Equation{A: 1, B: -4, C: 4},
		Dsc: 0,
		Str: "x*x - 4*x + 4 = 0",
		Rts: []float64{2},
	},
	// Случай, когда средний коэффициент равен 0
	{
		Equ: Equation{A: -1, B: 0, C: 1},
		Dsc: 4,
		Str: "-x*x + 1 = 0",
		Rts: []float64{1, -1},
	},
	// Случай, когда свободный коэффициент равен 0
	{
		Equ: Equation{A: 2, B: 4, C: 0},
		Dsc: 16,
		Str: "2*x*x + 4*x = 0",
		Rts: []float64{-2, 0},
	},
}

func TestGetDiscriminant(t *testing.T) {
	for i := range cases {
		cs := cases[i]
		res := cs.Equ.GetDiscriminant()

		if res != cs.Dsc {
			t.Error(
				"Случай:", cs.Equ,
				"Ожидается:", cs.Dsc,
				"Получено:", res,
			)
		}
	}
}

func TestToString(t *testing.T) {
	for i := range cases {
		cs := cases[i]
		res := cs.Equ.ToString()

		if res != cs.Str {
			t.Error(
				"Случай:", cs.Equ,
				"Ожидается:", cs.Str,
				"Получено:", res,
			)
		}
	}
}

func TestGetRoots(t *testing.T) {
	for iCs := range cases {
		var pass = true
		var cs = cases[iCs]
		var res, _ = cs.Equ.GetRoots()

		// Сравним корни
		if len(res) != len(cs.Rts) {
			pass = false
		} else {
			for iRt := range res {
				if cs.Rts[iRt] != res[iRt] {
					pass = false
					break
				}
			}
		}

		if !pass {
			t.Error(
				"Случай:", cs.Equ,
				"Ожидается:", cs.Rts,
				"Получено:", res,
			)
		}
	}
}
