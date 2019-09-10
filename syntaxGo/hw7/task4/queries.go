package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var urls = os.Args[1:]
	var finish chan *testURL
	var wg sync.WaitGroup

	// Если без аргументов, то читаем стандартный ввод.
	if len(urls) == 0 {
		scaner := bufio.NewScanner(os.Stdin)

		for scaner.Scan() {
			if line := scaner.Text(); len(line) != 0 && line != "\n" {
				urls = append(urls, line)
			} else {
				break
			}
		}
	}

	// Если нечего тестировать.
	if len(urls) == 0 {
		fmt.Println("Не выбраны URL")
		return
	}

	// Подготовка перед выполнением.
	finish = make(chan *testURL, len(urls))
	wg.Add(len(urls))
	fmt.Println()

	// Запустим параллельное выполнение.
	for i := range urls {
		go func(t *testURL) {
			defer wg.Done()
			t.Run(finish)
		}(&testURL{URL: urls[i]})
	}

	// Дождемся окончания выполнения всех потоков.
	wg.Wait()
	close(finish)

	// Победитель.
	finisher := <-finish
	fmt.Println("=====================================")
	fmt.Printf("Быстрее всех: %v\n", finisher.URL)
	fmt.Println("=====================================")
	fmt.Println()
	finisher.Print()

	// Остальные финишеры.
	for finisher = range finish {
		fmt.Println()
		finisher.Print()
	}

	fmt.Println()
}

// Описывает тестовый запрос по URL-адресу
type testURL struct {
	URL    string
	Status string
	Time   time.Duration
}

// Запустит тест
func (t *testURL) Run(finish chan<- *testURL) {
	var resp *http.Response
	var err error
	var t0, t1 time.Time

	// Выполним запрос
	t0 = time.Now()
	resp, err = http.Get(t.URL)
	t1 = time.Now()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Сохраним нужные части ответа
	t.Status = resp.Status
	t.Time = t1.Sub(t0)

	// Если получен статус об успешном запросе
	if resp.StatusCode == 200 {
		finish <- t
	}
}

// Выведет данные теста в консоль
func (t *testURL) Print() {
	fmt.Printf("   URL: %v\n", t.URL)
	fmt.Printf("Status: %v\n", t.Status)
	fmt.Printf("  Time: %v\n", t.Time)
}
