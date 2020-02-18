package phonebook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"

	pt "github.com/tatsushid/go-prettytable"
)

// Book - описывает телефонную книгу.
type Book struct {
	*Phones
}

// Load - загрузит список номеров из файла json.
func (b *Book) Load(filePath string) (err error) {
	var book = *b
	var file *os.File
	var fileData []byte

	// Откроем файл.
	file, err = os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	// Прочтем файл целиком.
	fileData, err = ioutil.ReadAll(file)
	if err != nil {
		return
	}

	// Преобразуем данные файла в структуру.
	err = json.Unmarshal(fileData, book.Phones)
	if err != nil {
		return
	}

	return
}

// Sort - отсортирует записи в телефонной книге.
func (b *Book) Sort() {
	sort.Sort(b)
}

// Print - выведет список записей в консоль в виде таблицы.
func (b *Book) Print() (err error) {
	var book = *b
	var phones = *(book.Phones)
	var table *pt.Table
	var columns = []pt.Column{
		{Header: "Имя", MinWidth: 20, MaxWidth: 40},
		{Header: "Номер", MinWidth: 12},
	}

	table, err = pt.NewTable(columns...)
	if err != nil {
		return err
	}
	table.Separator = " | "

	if phones == nil {
		return fmt.Errorf("Нет списка записей")
	}

	for i := range phones {
		phone := phones[i]
		name := fmt.Sprintf("%s %s", phone.LastName, phone.FirstName)
		table.AddRow(name, phone.Number)
	}

	table.Print()
	return
}
