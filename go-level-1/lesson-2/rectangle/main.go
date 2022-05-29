package main

import (
	"fmt"
)

func main() {
	var width = askForValue("Укажите ширину прямоугольника: ")
	var height = askForValue("Укажите высоту прямоугольника: ")
	fmt.Println("Площать прямоугольника:", width*height)
}

// Попросить значение у пользователя.
// Функция будет повторно запрашивать значение у пользователя до тех пор,
// пока тот не введет число в корректном виде. Например некорректным
// считается ввод букв или отрицательных чисел.
func askForValue(label string) float64 {
	var value float64
	var err error
	var errLabel = "Некорректный ввод:"

	for value <= 0 {
		fmt.Print(label)

		if value, err = scanValue(); err != nil {
			fmt.Println(errLabel, err)
			value = 0
		} else if value < 0 {
			fmt.Println(errLabel, "отрицательное значение")
			value = 0
		}
	}

	return value
}

// Считать введенное число из консоли.
// В случае, если неудастся считать число типа float, тогда вернется ошибка,
// а весь ввод будет очищен. Поэтому нельзя вводить сразу два числа в консоль.
func scanValue() (float64, error) {
	var value float64
	var err, inputError error
	var count int

	// Если корректный ввоод, просто вернем значение.
	if count, err = fmt.Scanln(&value); err == nil {
		return value, nil
	}

	// Отладочная часть...
	// fmt.Println("COUNT:", count)
	// fmt.Println("ERROR:", err)
	// fmt.Println("VALUE:", value)

	// Сохраним ошибку.
	inputError = err

	// Очистим некорректный ввод.
	for !(err == nil && count == 0) {
		count, err = fmt.Scanln()
	}

	return 0, inputError
}
