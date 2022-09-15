package main

import (
	"fmt"
	"sync"
)

func main() {
	var value int
	var wgr = sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		wgr.Add(1)

		go func(i int) {
			value = i
			wgr.Done()
		}(i)
	}

	wgr.Wait()
	fmt.Println("Value:", value)
}
