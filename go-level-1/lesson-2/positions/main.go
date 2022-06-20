package main

import (
	"fmt"
)

func main() {
	var source = askForValue("Введите трехзначное целое число: ")
	var p0, p1, p2 int

	p0 = source % 10
	p1 = source % 100 / 10
	p2 = source % 1000 / 100

	fmt.Printf("В числе %v сотен, %v десятков, %v единиц", p2, p1, p0)
}

// Попросить числовое значение у пользователя.
// Функция будет повторно запрашивать значение у пользователя до тех пор,
// пока тот не введет число в корректном виде. Если пользователь введет
// отрицательное число, функция сама преобразует его в положительное.
func askForValue(label string) int {
	var value int
	var err error
	var errLabel = "Некорректный ввод:"

	for {
		fmt.Print(label)

		if value, err = scanValue(); err != nil {
			fmt.Println(errLabel, err)
			value = 0
		} else {
			break
		}
	}

	if value < 0 {
		value *= -1
	}

	return value
}

// Считать введенное число из консоли.
// В случае, если неудастся считать число типа float, тогда вернется ошибка,
// а весь ввод будет очищен. Поэтому нельзя вводить сразу два числа в консоль.
func scanValue() (int, error) {
	var value int
	var err, inputError error
	var count int

	if count, err = fmt.Scanln(&value); err == nil {
		return value, nil
	}

	inputError = err

	for !(err == nil && count == 0) {
		count, err = fmt.Scanln()
	}

	return 0, inputError
}
