package main

import (
	"fmt"

	"./queue"
)

func main() {
	var strQ = new(queue.StringQ)

	strQ.Add("First")
	strQ.Add("Second")
	strQ.Add("Third")
	strQ.Add("4th", "5th", "6th")
	fmt.Printf("Всего в очереди %v элементов.\n", strQ.Len())

	for {
		item, ok := strQ.Take()

		if !ok {
			fmt.Println("\nОбход очереди завершен.")
			fmt.Printf("Всего в очереди %v элементов.\n", strQ.Len())
			break
		}

		fmt.Printf("Элемент очереди: %v\n", item)
	}
}
