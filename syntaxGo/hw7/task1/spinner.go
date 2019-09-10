package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)

	time.Sleep(10 * time.Second)
	fmt.Printf("\rВот и все...         \n")
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("Кручу спиннер: %c\r", r)
			time.Sleep(delay)
		}
	}
}
