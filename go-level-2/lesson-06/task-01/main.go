package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	var mtx = sync.Mutex{}
	var wgr = sync.WaitGroup{}
	var res = 0

	for i := 1; i <= 5; i++ {
		wgr.Add(1)

		go func() {
			mtx.Lock()
			res++
			mtx.Unlock()
			wgr.Done()
		}()
	}

	wgr.Wait()
	fmt.Println("Result:", res)
}
