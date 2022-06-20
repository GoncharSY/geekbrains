package main

import (
	"fmt"
	"math"
)

func main() {
	var number = askForNumber()
	var primes = findPrimes(number)
	fmt.Printf("Список простых чисел до %v: %v\n", number, primes)
}

// Попросить у пользователя число.
func askForNumber() int {
	var num int
	fmt.Print("Введите целое число: ")
	fmt.Scanln(&num)
	return num
}

// Найти простые до указанного числа. Решето Эратосфена.
func findPrimes(number int) []int {
	var primes []int
	var composites = make([]bool, number+1)
	var last = int(math.Sqrt(float64(number) + 1))

	for i := 2; i <= last; i++ {
		if composites[i] == false {
			for j := i * i; j <= number; j += i {
				composites[j] = true
			}
		}
	}

	for i := range composites {
		if i > 1 && !composites[i] {
			primes = append(primes, i)
		}
	}

	return primes
}
