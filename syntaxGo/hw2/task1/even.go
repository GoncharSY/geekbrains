// Написать функцию, которая определяет, четное ли число.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var iNum int64
	var sNum string
	var err error

	for {
		// Получим строку от пользователя.
		fmt.Printf("Введите целое число:")
		fmt.Scanln(&sNum)

		// Завершим выполнение программы, если нужно.
		switch sNum {
		case "exit":
			return
		case "stop":
			return
		case "q":
			return
		}

		// Преобразуем строку в целое число.
		iNum, err = strconv.ParseInt(sNum, 10, 0)
		if err != nil {
			continue
		}

		// Проверим на четность/нечетность.
		if mod := iNum % 2; mod == 0 {
			fmt.Printf("Число %v - ЧЕТНОЕ.\n", iNum)
		} else {
			fmt.Printf("Число %v - НЕЧЕТНОЕ.\n", iNum)
		}
	}
}
