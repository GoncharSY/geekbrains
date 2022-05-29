package main

import "fmt"

func main() {
	var width, height float64

	fmt.Print("Укажите ширину прямоугольника: ")
	fmt.Scanln(&width)

	fmt.Print("Укажите высоту прямоугольника: ")
	fmt.Scanln(&height)

	fmt.Println("Площать прямоугольника:", width*height)
}
