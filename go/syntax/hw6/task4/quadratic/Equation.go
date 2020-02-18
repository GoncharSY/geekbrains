package quadratic

import (
	"fmt"
	"math"
)

// Equation - описывает квадратное уравнение.
type Equation struct {
	A float64
	B float64
	C float64
}

// GetDiscriminant - Вычислит и вернет значение дикриминанта.
func (e *Equation) GetDiscriminant() float64 {
	var eq = *e
	var a, b, c = eq.A, eq.B, eq.C

	return math.Pow(b, 2) - 4*a*c
}

// GetRoots - Вычислит и вернет корни уравнения, х1 и х2.
// В случае, когда нет действительных корней, вернется ошибка.
func (e *Equation) GetRoots() (x []float64, err error) {
	var eq = *e
	var D = eq.GetDiscriminant()

	// Нет корней
	if D < 0 {
		err = fmt.Errorf("Дискриминант меньше нуля: вещественных корней нет")
		return
	}

	x = make([]float64, 0)

	// Только один корень.
	if D == 0 {
		x = append(x, -eq.B/2*eq.A)
		return
	}

	x = append(x, (-eq.B-math.Sqrt(D))/(2*eq.A))
	x = append(x, (-eq.B+math.Sqrt(D))/(2*eq.A))
	return
}

// ToString - Вернет строку, как записывается квадратное уравнение.
func (e *Equation) ToString() string {
	var eq = *e
	var part1 = "x*x"
	var part2 = "x"
	var part3 = ""

	if eq.A == -1 {
		part1 = fmt.Sprintf("-%s", part1)
	} else if eq.A != 1 {
		part1 = fmt.Sprintf("%v*%s", eq.A, part1)
	}

	if eq.B == 0 {
		part2 = ""
	} else if eq.B == 1 {
		part2 = fmt.Sprintf(" + %s", part2)
	} else if eq.B == -1 {
		part2 = fmt.Sprintf(" - %s", part2)
	} else if eq.B > 0 {
		part2 = fmt.Sprintf(" + %v*%s", eq.B, part2)
	} else {
		part2 = fmt.Sprintf(" - %v*%s", (-1 * eq.B), part2)
	}

	if eq.C > 0 {
		part3 = fmt.Sprintf(" + %v", eq.C)
	} else if eq.C < 0 {
		part3 = fmt.Sprintf(" - %v", (-1 * eq.C))
	}

	return fmt.Sprintf("%s%s%s = 0", part1, part2, part3)
}
