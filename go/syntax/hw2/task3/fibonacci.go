// Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная от 0.
// Числа Фибоначчи определяются соотношениями F0=0, F1=1, Fn=Fn-1 + Fn-2.

package main

import "fmt"

func main() {
	// fmt.Println(fibonacci1(20)) // Recursion
	fibonacci2(100) // Loop
	// fibonacci3(0, 0, 0, 100) // Recursion
}

// Вычислить посредством рекурсии.
// NOTE: Очень нагружает процессор. Очевидно, что НЕ эффективно!
func fibonacci1(iNum int) int {
	if iNum == 0 {
		return 0
	} else if iNum == 1 {
		return 1
	} else {
		return fibonacci1(iNum-1) + fibonacci1(iNum-2)
	}
}

// Вычислить посредством цикла.
// NOTE: на 94-й итерации происходит переполнение.
func fibonacci2(iNum uint64) {
	var iMem = [2]uint64{0, 1}

	for i := uint64(0); i < iNum; i++ {
		if i < 2 {
			fmt.Printf("%v\n", i)
		} else {
			iMem[0], iMem[1] = iMem[1], iMem[0]+iMem[1]

			if iMem[0] > iMem[1] {
				fmt.Printf("Переполнение на %v-й итерации...\n", i)
				break
			}

			fmt.Printf("%v\n", iMem[1])
		}
	}

	fmt.Println()
	return
}

// Вычислить посредством рекурсии, идея №2.
// NOTE: на 94-й итерации происходит переполнение.
// - n1 - Fn-2
// - n2 - Fn-1
// - iter - Текущая итерация
// - max - Макс. число итераций
func fibonacci3(n1, n2, iter, max uint) {
	// Выход из рекурсии
	if iter > max {
		fmt.Printf("Выполнение завершено!\n")
		return
	} else if n1 > n2 {
		fmt.Printf("Переполнение на %v-й итерации...\n", iter)
		return
	}

	// Печать текучего значения.
	fmt.Printf("%v\n", n2)

	// Следующий уровень рекурсии.
	if n2 == 0 {
		fibonacci3(n2, n1+n2+1, iter+1, max)
	} else {
		fibonacci3(n2, n1+n2, iter+1, max)
	}
}
