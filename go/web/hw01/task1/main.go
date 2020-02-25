package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func main() {
	var keyString = "<script>"
	var links = []string{
		"https://yandex.ru",
		"https://google.com",
		"https://rambler.ru",
		"https://yahoo.com",
		"https://aaaaabbbbbccccc.com",
		"https://xxxxxyyyyyzzzzz",
	}

	fmt.Println("Ключевая строка:", keyString)

	if results, err := findKey(keyString, links); err != nil {
		fmt.Println(err)
	} else if len(results) == 0 {
		fmt.Println("Ни по одной из ссылок ключевая строка не найдена.")
	} else {
		fmt.Println("Ключевая строка найдена по следующим адресам:")
		for i := range results {
			fmt.Println("-", results[i])
		}
	}
}

// Искать ключевую строку по адресам.
func findKey(key string, links []string) (results []string, err error) {
	var wgr sync.WaitGroup
	var mux sync.Mutex

	if len(key) == 0 {
		err = errors.New("Ключевая строка пуста")
		return
	}

	if len(links) == 0 {
		fmt.Println("Список ссылок пуст")
	}

	wgr.Add(len(links))

	for idx := range links {
		go func(link string) {
			defer wgr.Done()

			var err error
			var res *http.Response
			var body []byte

			// Запросим данные.
			if res, err = http.Get(link); err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			// Прочтем ответ.
			if body, err = ioutil.ReadAll(res.Body); err != nil {
				fmt.Println(err)
				return
			}

			// Поищем ключевую строку.
			if strings.Contains(string(body), key) {
				mux.Lock()
				results = append(results, link)
				mux.Unlock()
			}
		}(links[idx])
	}

	wgr.Wait()
	return
}
