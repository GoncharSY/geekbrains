package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var phoneBook, newPhones map[string][]int

	loadPhones(&phoneBook, "phoneBook.json")
	loadPhones(&newPhones, "newPhones.json")
	combineBooks(&phoneBook, &newPhones)
	saveBook(&phoneBook, "newBook.json")

	fmt.Println()
	for name, numbers := range phoneBook {
		fmt.Println("Абонент:", name)
		for i, number := range numbers {
			fmt.Printf("\t %v) %v \n", i+1, number)
		}
	}
}

// Сохранить данные телефонного справочника в файл.
func saveBook(book *map[string][]int, filePath string) {
	var fileData []byte
	var err error

	// Преобразуем данные в массив байт.
	fileData, err = json.Marshal(*book)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Запишем данные в файл.
	err = ioutil.WriteFile(filePath, fileData, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// Объединить справочники телефонных номеров.
func combineBooks(oldBook, newBook *map[string][]int) {
	var oB = *oldBook
	var nB = *newBook

	for name := range nB {
		if _, ok := oB[name]; ok {
			//TODO: Здесь хорошо бы сделать проверку на уникальность.
			oB[name] = append(oB[name], nB[name]...)
			continue
		}

		oB[name] = nB[name]
	}
}

// Загрузить телефонный справочник из файла.
func loadPhones(phoneBook *map[string][]int, filePath string) {
	var fileData []byte
	var file *os.File
	var err error

	// Откроем файл.
	file, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		file.Close()
		fmt.Println("ФАЙЛ", filePath, "ЗАКРЫТ")
	}()
	fmt.Println("ФАЙЛ", filePath, "ОТКРЫТ")

	// Прочтем файл целиком.
	fileData, err = ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Преобразуем данные файла в отображение.
	err = json.Unmarshal(fileData, phoneBook)
	if err != nil {
		log.Fatal(err)
		return
	}
}
