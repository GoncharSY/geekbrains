package main

import (
	"fmt"
	"time"
)

func main() {
	var cTick = time.Tick(1 * time.Second)
	var cStop = make(chan bool)
	var cStart = make(chan bool)

	fmt.Println("Обратный отсчет:")

	go scanComand(cStop)
	go printTimer(10, cTick, cStop, cStart)

	if <-cStart {
		fmt.Println("Пуск!")
	} else {
		fmt.Println("Запуск отменен.")
	}

	close(cStart)
	close(cStop)
}

// Сканирует вводимые пользователем команды циклично.
func scanComand(stop chan<- bool) {
	var comand string

	for {
		fmt.Scan(&comand)

		switch comand {
		case "exit":
			stop <- true
		default:
			continue
		}
	}
}

// Отобразит таймер в консоли.
func printTimer(steps int, tick <-chan time.Time, stop <-chan bool, start chan<- bool) {
	for itr := steps; itr > 0; itr-- {
		select {
		case <-stop:
			start <- false
			return
		case t := <-tick:
			fmt.Printf("%v == %v ==\n", t.Format("15:04:05"), itr)
		}
	}

	start <- true
}
