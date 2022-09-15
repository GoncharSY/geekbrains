// Это код программы для третьего задания.
// - Удаляется существующая папка с нужным именем.
// - Создается новая пустая папка с нужным именем.
// - В папку добавляется 1 млн. пустых файлов.
// - Функция, создающая файл, закрывает файл при завершении.
// - В процессе создания выводятся сообщения в консоль.

package main

import (
	"errors"
	"fmt"
	"os"
)

const folderName = "tmp"
const fileCount = 10
const namePrefix = "File-"

// const printStep = fileCount/10 - 1

func main() {
	fmt.Println("Files creating...")

	if err := createFolder(folderName); err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < fileCount; i++ {
		name := fmt.Sprintf("%s/%s%v", folderName, namePrefix, i)

		if err := createFile(name); err != nil {
			fmt.Println(fmt.Errorf("error creating file '%s': %w", name, err))
			return
		}

		if (i+1)%10 == 0 {
			fmt.Println("   ", i+1, "files created")
		}
	}

	fmt.Println("Creating finished.")
}

// Создать новую папку для файлов.
// Перед созданием новой папки сначала будет удалена существующая с тем же именем.
func createFolder(name string) error {
	if name == "" {
		return errors.New("empty name")
	}
	if err := os.RemoveAll(name); err != nil {
		return fmt.Errorf("removing old folder: %w", err)
	}
	if err := os.Mkdir(name, os.ModeDir); err != nil {
		return fmt.Errorf("creating new folder: %w", err)
	}

	return nil
}

// Создать новый файл.
func createFile(name string) error {
	if name == "" {
		return errors.New("empty name")
	}

	var file *os.File
	var err error

	if file, err = os.Create(name); err == nil {
		return err
	}

	defer file.Close()
	return err
}
