// Это код программы для первого задания.
// - Паническая ситуация создается.
// - Паника перехватывается.
// - Сообщение в консоль выводится.

package main

import "fmt"

func main() {
	makeImpPanic()
	fmt.Println("run complete")
}

// Создает неявную панику, обращаясь к несущствующему элементу среза.
// Затем отложенным вызовом отлавливает паническую ситуацию и сообщает об этом в консоль.
func makeImpPanic() {
	var a = make([]int, 1)

	defer recoverImpPanic()

	a[1] = 7
}

// Отлавливает паническую ситуацию, созданную неявно.
func recoverImpPanic() {
	if err := recover(); err != nil {
		fmt.Println("panic is recovered:", err)
	}
}
