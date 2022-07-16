package fibonacci

// GetCalculator returns function for calculating number from Fibonacci row by index.
func GetCalculator(optimised bool) func(int) (int, error) {
	if !optimised {
		return CalculateNumber
	}

	return func(n int) (int, error) {
		return CalculateWithBuffer(n, nil)
	}
}
