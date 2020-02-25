package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

// RequesBody - описывает структуру тела запроса.
type RequesBody struct {
	Key   string   `json:"search"`
	Links []string `json:"sites"`
}

func main() {
	var srvMux *http.ServeMux

	srvMux = http.NewServeMux()
	srvMux.HandleFunc("/", onRequest)

	fmt.Println("Запуск сервера...")
	log.Fatal(http.ListenAndServe(":8080", srvMux))
}

// Обработать запрос.
func onRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Println()
	fmt.Println("Пришел запрос:", req.URL.Path)
	fmt.Println("Метод запроса:", req.Method)

	res.Header().Set("Access-Control-Allow-Origin", "*")

	if req.Method == "GET" {
		res.WriteHeader(200)
		res.Write([]byte("Используй POST-запрос!"))
		return
	} else if req.Method != "POST" {
		res.WriteHeader(404)
		res.Write([]byte(""))
		return
	}

	var (
		err     error
		okLinks = make([]string, 0)
		body    = make([]byte, 0)
		resBody = make([]byte, 0)
		reqBody = RequesBody{}
	)

	// Получим тело запроса.
	if body, err = ioutil.ReadAll(req.Body); err != nil {
		res.WriteHeader(500)
		res.Write([]byte("Не удалось прочесть тело запроса:" + err.Error()))
		return
	} else if len(body) == 0 {
		res.WriteHeader(200)
		res.Write([]byte("Тело запроса пусто"))
		return
	}

	// Разберем тело запроса.
	if err := json.Unmarshal(body, &reqBody); err != nil {
		res.WriteHeader(500)
		res.Write([]byte([]byte("Не удалось разобрать JSON: " + err.Error())))
		return
	}

	// Выполним поиск ключа по ссылкам.
	if okLinks, err = findKey(reqBody.Key, reqBody.Links); err != nil {
		res.WriteHeader(200)
		res.Write([]byte(err.Error()))
		return
	}

	// Подготовим ответ.
	if resBody, err = json.Marshal(okLinks); err != nil {
		res.WriteHeader(500)
		res.Write([]byte("Не удалось создать тело ответа: " + err.Error()))
		return
	}

	res.WriteHeader(200)
	res.Write(resBody)
}

// Искать ключевую строку по адресам.
func findKey(key string, links []string) (results []string, err error) {
	var (
		wgr sync.WaitGroup
		mux sync.Mutex
	)

	if len(key) == 0 {
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
			defer wgr.Done()

			var (
				err  error
				res  *http.Response
				body = make([]byte, 0)
			)

			if res, err = http.Get(link); err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			if body, err = ioutil.ReadAll(res.Body); err != nil {
				fmt.Println(err)
				return
			}

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
