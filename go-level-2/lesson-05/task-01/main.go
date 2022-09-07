package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var limit = 5
	var strGrp = sync.WaitGroup{}
	var stpGrp = sync.WaitGroup{}

	for i := 1; i <= limit; i++ {
		strGrp.Add(1)
		stpGrp.Add(1)
		go runGR(i, &strGrp, &stpGrp)
	}

	strGrp.Wait()
	fmt.Printf("All goroutenes (%d) have been STARTED\n", limit)

	stpGrp.Wait()
	fmt.Printf("All goroutenes (%d) have been STOPPED\n", limit)
}

func runGR(num int, str, stp *sync.WaitGroup) {
	fmt.Printf("Goroutine-%d: started\n", num)
	str.Done()

	time.Sleep(1 * time.Second)

	fmt.Printf("Goroutine-%d: stopped\n", num)
	stp.Done()
}
