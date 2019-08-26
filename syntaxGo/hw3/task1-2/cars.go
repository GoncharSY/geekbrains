package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"./cars"
)

func main() {
	var filePath = "cars.json"
	var fileData []byte
	var file *os.File
	var carPark []cars.Car
	var err error

	// Откроем файл.
	file, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		file.Close()
		fmt.Println("ФАЙЛ ЗАКРЫТ")
	}()
	fmt.Println("ФАЙЛ ОТКРЫТ")

	// Прочтем файл целиком.
	fileData, err = ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Преобразуем данные файла в структуру.
	err = json.Unmarshal(fileData, &carPark)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Выведем результат.
	for i := range carPark {
		fmt.Println("=============================================")
		printCarInfo(carPark[i])
		fmt.Println("=============================================")
	}
}

// Вывести описание автомобиля в консоль.
func printCarInfo(car cars.Car) {
	fmt.Printf("             Модель: %v\n", car.Model)
	fmt.Printf("                Тип: %v\n", car.Type)
	fmt.Printf("        Год выпуска: %v\n", car.Release.Year())
	fmt.Println("                  --")
	printEngineInfo(car.Engine)
	fmt.Println("                  --")
	printTrunkInfo(car.Trunk)
	fmt.Println("                  --")
	printWindowsInfo(car.Windows)
}

// Вывести описание двигателя автомобиля.
func printEngineInfo(e cars.Engine) {
	var unitCapacity = "л.с."
	var state = map[bool]string{
		true:  "Работает",
		false: "Не работает",
	}

	fmt.Printf(" Мощность двигателя: %v %v\n", e.Capacity, unitCapacity)
	fmt.Printf("Состояние двигателя: %v\n", state[e.IsRunning])
}

// Вывести описание багажника/кузова автомобиля.
func printTrunkInfo(t cars.Trunk) {
	var unit = "л"

	fmt.Printf("    Объем багажника: %v%v\n", t.Volume, unit)
	fmt.Printf("        Заполнен на: %v%v (%v%%)\n", t.Volume/100*t.Filled, unit, t.Filled)
}

// Вывести описание окон автомобиля.
func printWindowsInfo(windows []cars.Window) {
	var state = map[bool]string{
		true:  "Открыто",
		false: "Закрыто",
	}

	// Все окна.
	fmt.Printf("    Количество окон: %v\n", len(windows))

	// Каждое отдельное окно.
	for i := range windows {
		fmt.Printf("           Окно №%v: %v\n", i+1, state[windows[i].IsOpen])
	}
}
