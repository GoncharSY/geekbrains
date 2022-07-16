package config

import (
	"fmt"
	"go-level-1/lesson-10/fibonacci"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	var cfg = New()

	assert.Equal(t, true, cfg.UseBuffer, "test default value of buffer mode")
	assert.Equal(t, 0, cfg.Index, "test default index of number")

	cfg.UseBuffer = false
	cfg.Index = 17

	assert.Equal(t, false, cfg.UseBuffer, "test changed value of buffer mode")
	assert.Equal(t, 17, cfg.Index, "test changed index of number")
}

func ExampleNew() {
	var cfg = New()
	var idx int = cfg.Index
	var fib = fibonacci.GetCalculator(cfg.UseBuffer)

	if cfg.UseBuffer {
		fmt.Println("Optimised calculation is enabled")
	} else {
		fmt.Println("Optimised calculation is disabled")
	}

	if num, err := fib(idx); err != nil {
		panic(fmt.Errorf("error of calculation for %v: %w", idx, err))
	} else {
		fmt.Printf("Fibonacci number with index %v: %v\n", idx, num)
	}
}
