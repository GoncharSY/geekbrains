package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var slice = make([]int, 10)

	fillSlice(slice)
	fmt.Println("Исходный набор:", slice)

	sortSlice(slice)
	fmt.Println("Сортированный набор:", slice)
}

// Заполнить срез целыми числами.
func fillSlice(slice []int) {
	rand.Seed(time.Now().UnixNano())

	for i := range slice {
		slice[i] = rand.Intn(10)
	}
}

// Отсортировать набор чисел методом вставок.
func sortSlice(slice []int) {
	for i := 1; i < len(slice); i++ {
		for j := i - 1; j >= 0 && slice[j] > slice[j+1]; j-- {
			slice[j], slice[j+1] = slice[j+1], slice[j]
		}
	}
}

// ПСЕВДОКОД ИЗ ВИКИПЕДИИ
// for j = 2 to A.length do
//     key = A[j]
//     i = j-1
//     while (int i > 0 and A[i] > key) do
//         A[i + 1] = A[i]
//         i = i - 1
//     end while
//     A[i+1] = key
// end

// КОД НА ОСНОВЕ ПСЕВДОКОДА
// var item int
// var i, j int

// for i = 1; i < len(slice); i++ {
// 	item = slice[i]
// 	j = i - 1

// 	for j >= 0 && slice[j] > item {
// 		slice[j+1] = slice[j]
// 		j--
// 	}

// 	slice[j+1] = item
// }
