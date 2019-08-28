package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var progName = os.Args[0]
var file = flag.String("f", "", "Файл, в котором будем искать подходящий текст.")
var expr = flag.Bool("e", false, "Если включено, то строка поиска воспринимается, как регулярное выражение.")
var args []string

func main() {
	flag.Parse()
	args = flag.Args()

	// Отсутствие параметров, как зов о помощи.
	if len(args) == 0 || args[0] == "" {
		printHelp()
		return
	}

	var file = *file
	var expr = *expr
	var match = getFindFunction(expr, args[0])

	// Текущие параметры.
	printOpts()

	// Определим место поиска (стандартный ввод или файл).
	if file == "" {
		findInStd(match)
	} else {
		findInFile(file, match)
	}
}

// printHelp - выведет справочную информацию для пользователя.
func printHelp() {
	fmt.Println("Программа для поиска текстовых строк, содержащих фрагмент или удовлетворяющих регулярному выражению.")
	fmt.Println("Если передать несколько строк для поиска, то поиск будет выполняться только по первой из них.")
	fmt.Println("Пример:")
	fmt.Printf("  %s 'search text'\n", progName)
	fmt.Printf("  %s -e 'regular expression'\n", progName)
	fmt.Printf("  %s -f='file.ext' 'search text'\n", progName)
	fmt.Printf("  %s -e -f='file.ext' 'regular expression'\n", progName)
	fmt.Println("Параметры:")
	flag.PrintDefaults()
}

// printOpts - выведет опции с которыми была запущена программа.
func printOpts() {
	if *file != "" {
		fmt.Printf("Исходный файл: %s\n", *file)
	}

	if *expr {
		fmt.Printf("   Тип поиска: регулярное выражение\n")
		fmt.Printf("Рег.выражение: %v\n", args[0])
	} else {
		fmt.Printf("   Тип поиска: подстрока\n")
		fmt.Printf("    Подстрока: %v\n", args[0])
	}

	fmt.Println()
}

// getFindFunction - вернет нужную функцию поиска.
// Поиск будет выполняться либо по рег.выражению, либо по подстроке.
func getFindFunction(regExp bool, key string) func(s string) bool {
	if regExp {
		return func(str string) bool {
			matched, err := regexp.Match(key, []byte(str))
			if err != nil {
				fmt.Println(err)
			}

			return matched
		}
	} else {
		return func(str string) bool {
			return strings.Contains(str, key)
		}
	}
}

// findInStd - будет искать по тексту из стандартного потока.
func findInStd(match func(s string) bool) {
	var scaner = bufio.NewScanner(os.Stdin)

	// Построчное чтение.
	for scaner.Scan() {
		line := scaner.Text()

		if match(line) {
			fmt.Println(line)
		}
	}

	// Проверим наличие ошибок в конце.
	if err := scaner.Err(); err != nil {
		log.Fatal(err)
	}
}

// findInFile - будет искать по тексту из файла.
func findInFile(path string, match func(s string) bool) {
	var file *os.File
	var scaner *bufio.Scanner
	var err error

	// Откроем файл.
	file, err = os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Читаем файл построчно.
	scaner = bufio.NewScanner(file)
	for scaner.Scan() {
		line := scaner.Text()

		if match(line) {
			fmt.Println(line)
		}
	}

	// Проверим наличие ошибок в конце.
	if err := scaner.Err(); err != nil {
		log.Fatal(err)
	}
}
