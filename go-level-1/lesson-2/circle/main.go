package main

import (
	"fmt"
	"math"
)

func main() {
	var radius, diameter, length float64
	var area = askForValue("Введите площать круга: ")

	radius = math.Sqrt(area / math.Pi)
	diameter = 2 * radius
	length = diameter * math.Pi

	fmt.Println("Диаметр круга:", diameter)
	fmt.Println("Длина окружности круга:", length)
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

	if count, err = fmt.Scanln(&value); err == nil {
		return value, nil
	}

	inputError = err

	for !(err == nil && count == 0) {
		count, err = fmt.Scanln()
	}

	return 0, inputError
}
