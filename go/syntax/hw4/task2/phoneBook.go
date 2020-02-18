package main

import (
	"fmt"

	pb "./phonebook"
)

func main() {
	var book = &pb.Book{Phones: new(pb.Phones)}
	var err = book.Load("phoneBook.json")

	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("\n\n-- До сортировки:")
	err = book.Print()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("\n\n-- После сортировки: Фамилия-Имя-Номер")
	book.Sort()
	err = book.Print()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
}
