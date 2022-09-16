package main

// Человек.
type Human struct {
	Name   string
	Age    uint
	Height centimeter
	Weight kilogram
}

// Мера роста человека (сантиметр).
type centimeter float64

// Мера массы человека (килограмм).
type kilogram float32
