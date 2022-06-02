package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var op = askForOperator()
	var a, b = askForOperands(op)
	var res float64

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка деления на 0 (ноль)")
			os.Exit(1)
		}
		res = a / b
	case "pow":
		res = math.Pow(a, b)
	case "fact":
		res = float64(calculateFact(int(a)))
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}

// Попросить у пользователя оператор.
func askForOperator() string {
	var op string

	fmt.Print("Введите арифметическую операцию (+, -, *, /, pow, fact): ")
	fmt.Scanln(&op)
	return op
}

// Попросить у пользователя операнды.
func askForOperands(op string) (float64, float64) {
	var a, b float64

	switch op {
	case "fact":
		fmt.Print("Введите число: ")
		fmt.Scanln(&a)
	default:
		fmt.Print("Введите первое число: ")
		fmt.Scanln(&a)
		fmt.Print("Введите второе число: ")
		fmt.Scanln(&b)
	}

	return a, b
}

// Посчитать факториал числа
func calculateFact(a int) int {
	var res int = 1

	if a != 0 {
		for i := 1; i <= a; i++ {
			res *= i
		}
	}

	return res
}
