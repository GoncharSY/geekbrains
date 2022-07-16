package main

import (
	"fmt"
	"go-level-1/lesson-10/config"
	"go-level-1/lesson-10/fibonacci"
)

func main() {
	var cfg = config.New()
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
