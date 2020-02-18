package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var progName = os.Args[0]
var origin = flag.String("file", "", "Имя оригинального файла.")
var copy = flag.String("copy", "", "Имя нового файла-копии.")

func main() {
	flag.Parse()

	// Проверим наличие обязательных параметров.
	if *origin == "" || *copy == "" {
		printHelp()
		return
	}

	fmt.Printf("Исходный файл: %s\n", *origin)
	fmt.Printf("        Копия: %s\n", *copy)
	fmt.Println()

	var fileData []byte
	var file *os.File
	var err error

	// Откроем файл.
	file, err = os.Open(*origin)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Прочтем файл целиком.
	fileData, err = ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Запишем данные в новый файл.
	if _, err := os.Stat(*copy); err != nil {
		if os.IsNotExist(err) {
			err = ioutil.WriteFile(*copy, fileData, 0644)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
		return
	}

	// Перезапишем существующий файл.
	if getConfirm(fmt.Sprintf("Файл '%s' уже существует. Желаете перезаписать этот файл?", *copy)) {
		err = ioutil.WriteFile(*copy, fileData, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// printHelp - выведет справочную информацию для пользователя.
func printHelp() {
	fmt.Println("Программа для копирования файла.")
	fmt.Println("Пример:")
	fmt.Printf("  %s -file=<file-name.ext> -copy=<copy-name.ext>\n", progName)
	fmt.Println("Параметры:")
	flag.PrintDefaults()
}

// getUserConfirm - запросит подтверждение у пользователя и вернет его результат.
func getConfirm(text string) bool {
	var scaner = bufio.NewScanner(os.Stdin)
	var format = "%s (y/n): "

	fmt.Printf(format, text)

	for scaner.Scan() {
		ans := scaner.Text()

		switch strings.ToLower(ans) {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Printf(format, text)
		}
	}

	if err := scaner.Err(); err != nil {
		log.Fatal(err)
	}

	return false
}
