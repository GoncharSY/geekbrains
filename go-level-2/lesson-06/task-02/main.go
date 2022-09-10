package main

import (
	"fmt"
	"runtime"
)

func main() {
	// runtime.GOMAXPROCS(1)

	for i := 0; i < 5; i++ {
		go fmt.Println("Goroutine-", i)
	}

	for i := 0; i < 500e6; i++ {
		if i%1e6 == 0 {
			runtime.Gosched()
		}
	}

	fmt.Println("Done")
}
