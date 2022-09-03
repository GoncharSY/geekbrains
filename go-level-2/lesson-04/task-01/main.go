package main

import (
	"fmt"
)

const buffer = 100 // Максимум запущенных горутин.
const limit = 1000 // Общее число горутин.

func main() {
	var result = 0
	var buff = make(chan struct{}, buffer)
	var pool = make(chan struct{}, limit)

	// Запуск горутин.
	for i := 0; i < limit; i++ {
		buff <- struct{}{}

		go func() {
			defer func() {
				pool <- <-buff
			}()

			result++
		}()
	}

	// Завешение горутин.
	for i := 0; i < limit; i++ {
		<-pool
	}

	fmt.Println("Result number:", result)
}
