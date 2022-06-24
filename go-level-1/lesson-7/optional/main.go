/*
	Необходимо объявить свой тип, обернув в него тип `[]byte` - (слайс байтов).

	Затем, необходимо реализовать на нем такие методы, чтобы он удовлетворял
	интферйесам `io.Reader` (из него можно читать байты) и `io.Writer` (а так же
	писать их туда). Затем, используя функции пакета `io`:
	- С помощью `io.WriteString` записать в переменную вашего типа произвольную
	  строку.
	- С помощью `io.ReadAll( )` считать вашу строку обратно (вообще говоря, он
	  возвращает слайс байт, но его легко привести к виду строки)

	В случае затруднений можно писать мне, обсудим (tg @VladimirLozhkin), или
	обратиться к коду реализации типа `Buffer{}` из пакета `bytes`. (Но там
	спойлеры, сначала попробуйте сами :D )
*/

package main

import (
	"fmt"
	"io"
)

// Хранилище данных.
type DataKeeper struct {
	data []byte
}

// Считает данные из хранилища в срез байт.
func (this DataKeeper) Read(p []byte) (int, error) {
	var length int = len(p)
	var err error = nil

	for i := range p {
		if i >= len(this.data) {
			length = len(this.data)
			err = io.EOF
			break
		}

		p[i] = this.data[i]
	}

	return length, err
}

// Запишет данные из среза байт в хранилище.
func (this *DataKeeper) Write(p []byte) (n int, err error) {
	this.data = make([]byte, len(p))

	for i := range p {
		this.data[i] = p[i]
	}

	return len(p), nil
}

func main() {
	var wd = []byte("Some simple and short string")
	var dk = &DataKeeper{}
	var rd10 = make([]byte, 10)
	var rd50 = make([]byte, 50)
	var rd []byte
	var err error
	var num int

	// Запись данных в хранилище.
	if num, err = dk.Write(wd); err != nil {
		fmt.Println(fmt.Errorf("Error while writting: %w", err))
	} else {
		fmt.Printf("Writting %v bytes\n", num)
		fmt.Printf("Data: %v\n---\n", *dk)
	}

	// Чтение части данных.
	// В данном случае срез, куда я считываю, меньше чем объем данных в хранилище.
	// Поэтому часть данных просто не попадет в целевой срез.
	switch num, err = dk.Read(rd10); err {
	case io.EOF:
		fmt.Println("Data read to the end")
		fallthrough
	case nil:
		fmt.Printf("Read %v bytes\n", num)
		fmt.Printf("Data: %v\n---\n", string(rd10))
	default:
		fmt.Println(fmt.Errorf("Error while reading: %w", err))
	}

	// Чтение данных целиком.
	// В данном случае срез, куда я считываю, больше чем объем данных в хранилище.
	// Поэтому все данные из хранилища попадут в целевой срез.
	switch num, err = dk.Read(rd50); err {
	case io.EOF:
		fmt.Println("Data read to the end")
		fallthrough
	case nil:
		fmt.Printf("Read %v bytes\n", num)
		fmt.Printf("Data: %v\n---\n", string(rd50))
	default:
		fmt.Println(fmt.Errorf("Error while reading: %w", err))
	}

	// Запись строки в хранилище.
	if num, err = io.WriteString(dk, string(wd)); err != nil {
		fmt.Println(fmt.Errorf("Error while writting of a string: %w", err))
	} else {
		fmt.Printf("Writting by io.WriteString (%v bytes)\n", num)
		fmt.Printf("Data: %v\n---\n", *dk)
	}

	// Чтение данных целиком.
	switch rd, err = io.ReadAll(dk); err {
	case io.EOF:
		fmt.Println("Data read to the end:", err)
		fallthrough
	case nil:
		fmt.Println("All data was read by io.ReadAll without error.")
		fmt.Printf("Data: %v\n---\n", string(rd))
	default:
		fmt.Println(fmt.Errorf("Error while ReadAll: %w", err))
	}
}
