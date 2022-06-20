// ЗАДАЧА: Оптимизируйте приложение за счёт сохранения предыдущих результатов в мапе.

package main

import (
	"fmt"
)

var fibonacci = make(map[int]int)

func main() {
	var number = askForNumber()
	var fibonacci []int

	for i := 0; i <= number; i++ {
		fibonacci = append(fibonacci, getFibonacci(i))
	}

	fmt.Printf("Числа Фибоначчи от 0 до %v: %v\n", number, fibonacci)
}

// Задать исходное состояние программы.
func init() {
	fibonacci[0] = 0
	fibonacci[1] = 1
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
	if num, ok := fibonacci[n]; ok {
		return num
	}

	// Вывод для отладки вызовов функции.
	// fmt.Println("Вычисляю", n)

	fibonacci[n] = getFibonacci(n-1) + getFibonacci(n-2)
	return fibonacci[n]
}
