package main

import (
	"flag"
	"fmt"
	"os"

	"./quadratic"
)

var progName = os.Args[0]
var coefA float64
var coefB float64
var coefC float64

func main() {
	parseInput()

	if coefA == 0 {
		help()
		return
	}

	var quadEq = &quadratic.Equation{
		A: coefA,
		B: coefB,
		C: coefC,
	}

	fmt.Println()
	fmt.Println("Уравнение:", quadEq.ToString())
	fmt.Println("Дискриминант:", quadEq.GetDiscriminant())

	if roots, err := quadEq.GetRoots(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Корни уравнения:")
		for i, root := range roots {
			fmt.Printf("   x%v = %v\n", i+1, root)
		}
	}

	fmt.Println()
}

// Разберет входящие параметры.
func parseInput() {
	flag.Float64Var(&coefA, "a", 0.0, "Первый (старший) коэффициент.")
	flag.Float64Var(&coefB, "b", 0.0, "Второй (средний) коэффициент при x")
	flag.Float64Var(&coefC, "c", 0.0, "Свободный член")
	flag.Parse()
}

// Отобразит в консоли подсказку для пользователя.
func help() {
	fmt.Println()
	fmt.Println("Программа для вычисления коней квадратного уравнения.")
	fmt.Println("Помните, что старший коэффициент не должен равняться 0.")
	fmt.Println("\nПример:")
	fmt.Printf("   %s -a=1 -b=2 -c=3\n", progName)
	fmt.Println("\nПараметры:")
	flag.PrintDefaults()
	fmt.Println()
}
