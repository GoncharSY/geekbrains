package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
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
