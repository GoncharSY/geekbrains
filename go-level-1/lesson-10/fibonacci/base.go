package fibonacci

import (
	"errors"
)

// CalculateNumber calculates number from the Fibonacci row.
func CalculateNumber(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("positive numbers only")
	}

	if n < 2 {
		return n, nil
	}

	var n1, n2 int
	var err error

	if n2, err = CalculateNumber(n - 2); err != nil {
		return 0, err
	}
	if n1, err = CalculateNumber(n - 1); err != nil {
		return 0, err
	}

	return n1 + n2, nil
}
