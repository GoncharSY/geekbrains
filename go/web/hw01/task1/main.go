package main

import "fmt"
import "errors"
import "net/http"
import "io/ioutil"
import "strings"
import "sync"

func main() {
	var keyString = "<body"
	var links = []string{
		"https://yandex.ru",
		"https://google.com",
		"https://rambler.ru",
		"https://yahoo.com",
		"https://aaaaabbbbbccccc",
		"https://xxxxxyyyyyzzzzz",
	}

	fmt.Println()
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

	fmt.Println()
}

// Искать ключевую строку по адресам.
func findKey(key string, links []string) (results []string, err error) {
	var wgr sync.WaitGroup

	if key == "" {
		err = errors.New("Ключевая строка пуста")
		return
	}

	if len(links) == 0 {
		err = errors.New("Список ссылок пуст")
		return
	}

	wgr.Add(len(links))

	for idx := range links {
		go func(link string) {
			var res *http.Response
			var body []byte

			defer wgr.Done()

			// Запросим данные.
			if res, err = http.Get(link); err != nil {
				err = nil
				return
			}
			defer res.Body.Close()

			// Прочтем ответ.
			if body, err = ioutil.ReadAll(res.Body); err != nil {
				err = nil
				return
			}

			// Поищем ключевую строку.
			if strings.Contains(string(body), key) {
				results = append(results, link)
			}
		}(links[idx])
	}

	wgr.Wait()
	return
}
