package main

import (
	"fmt"
)

const buffer = 4   // Максимум запущенных горутин.
const limit = 1000 // Общее число горутин.

func main() {
	var buff = make(chan struct{}, buffer)
	var pool = make(chan struct{}, limit)
	var box = make(chan uint, 1)

	// Кладем в коробку.
	box <- 0

	// Запускаем горутины.
	for i := 0; i < limit; i++ {
		buff <- struct{}{}

		go func() {
			defer func() {
				pool <- <-buff
			}()

			box <- (<-box) + 1
		}()
	}

	// Ждем завершения.
	for i := 0; i < limit; i++ {
		<-pool
	}

	// Проверяем результат.
	fmt.Println("Result number:", <-box)
}
