package main

import (
	"fmt"
	"sync"
)

var cnt = 0
var mtx = sync.Mutex{}
var grp = sync.WaitGroup{}

func main() {
	for i := 1; i <= 1000; i++ {
		grp.Add(1)
		go addCounter()
	}

	grp.Wait()
	fmt.Printf("Result: %d\n", cnt)
}

func addCounter() {
	mtx.Lock()
	defer mtx.Unlock()

	cnt++
	grp.Done()
}
