package main

// Мера роста человека (сантиметр).
type centimeter float64

// Мера массы человека (килограмм).
type kilogram float32

// Человек.
type Human struct {
	Name   string
	Age    uint
	Height centimeter
	Weight kilogram
}

// IsEqualFields проверят равенство всех полей у двух объектов.
// Если все поля соответсвенно равны возвращается true иначе false.
func (h1 *Human) IsEqualFields(h2 *Human) bool {
	return h1.Name == h2.Name &&
		h1.Age == h2.Age &&
		h1.Height == h2.Height &&
		h1.Weight == h2.Weight
}
