package fibonacci

import (
	"errors"
)

// CalculateNumber calculates number from the Fibonacci row.
// You can give nil as the buffer. This function uses inner buffer while running calculation.
func CalculateWithBuffer(n int, buffer *map[int]int) (int, error) {
	if n < 0 {
		return 0, errors.New("positive numbers only")
	}

	if buffer == nil || *buffer == nil {
		buffer = &map[int]int{
			0: 0,
			1: 1,
		}
	}

	if number, ok := (*buffer)[n]; ok {
		return number, nil
	}

	var n1, n2 int
	var err error

	if n1, err = CalculateWithBuffer(n-1, buffer); err != nil {
		return 0, err
	}
	if n2, err = CalculateWithBuffer(n-2, buffer); err != nil {
		return 0, err
	}

	(*buffer)[n] = n1 + n2
	return (*buffer)[n], nil
}
