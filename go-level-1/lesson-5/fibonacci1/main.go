// ЗАДАЧА: Напишите приложение, рекурсивно вычисляющее заданное из стандартного ввода число Фибоначчи.

package main

import (
	"fmt"
)

func main() {
	var number = askForNumber()
	var fibonacci []int

	for i := 0; i <= number; i++ {
		fibonacci = append(fibonacci, getFibonacci(i))
	}

	fmt.Printf("Числа Фибоначчи от 0 до %v: %v\n", number, fibonacci)
}

// Попросить у пользователя число.
func askForNumber() int {
	var num int
	fmt.Print("Введите номер числа Фибоначчи: ")
	fmt.Scanln(&num)
	return num
}

// Получить N-e число фибоначи посредством рекусрии.
func getFibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	// Вывод для отладки вызовов функции.
	// fmt.Println("Вычисляю", n)

	return getFibonacci(n-1) + getFibonacci(n-2)
}
