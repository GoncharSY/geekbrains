/*
Перепишите программу-конвейер, ограничив количество передаваемых для обработки
значений и обеспечив корректное завершение всех горутин.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	iter := 10
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		defer func() {
			fmt.Println("Генерация завершена.")
			close(naturals)
		}()
		for x := 0; x < iter; x++ {
			naturals <- x
		}
	}()

	// возведение в квадрат
	go func() {
		defer func() {
			fmt.Println("Преобразование завершено.")
			close(squares)
		}()
		for x := range naturals {
			squares <- x * x
		}
	}()

	// печать
	defer func() {
		fmt.Println("Печать завершена.")
	}()
	for res := range squares {
		fmt.Println(res)
		time.Sleep(500 * time.Millisecond)
	}
}
