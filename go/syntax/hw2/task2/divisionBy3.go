// Написать функцию, которая определяет, делится ли число без остатка на 3.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var iDiv int64
	var fNum, fDiv float64
	var sNum string
	var err error

	for {
		// Получим строку от пользователя.
		fmt.Printf("Введите любое десятичное число:")
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

		// Преобразуем строку в число.
		fNum, err = strconv.ParseFloat(sNum, 64)
		if err != nil {
			continue
		}

		// Проверим наличие остатка при деление на 3.
		fDiv = fNum / 3
		iDiv = int64(fDiv)

		if fDiv == float64(iDiv) {
			fmt.Printf("Число %v - ДЕЛИТСЯ на 3 без остатка.\n", fNum)
		} else {
			fmt.Printf("Число %v - НЕ ДЕЛИТСЯ на 3 без остатка.\n", fNum)
		}
	}
}
