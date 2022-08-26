// Это код с решением четвертого практического задания.
// В данном примере паника создается и перехватывается в одной горутине.
// При перехвате паники в консоль печатается сообщение.
// Основаная программа завершается планово, через одну секунду.

package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered", v)
		}
	}()

	go func() {
		defer RecoverPanic()
		panic("A-A-A!!!")
	}()

	time.Sleep(time.Second)
	fmt.Println("execution completed")
}

// RecoverPanic перехватывает панику и печатает сообщение в стандартный поток вывода.
func RecoverPanic() {
	if err := recover(); err != nil {
		fmt.Println("recovered:", err)
	}
}
