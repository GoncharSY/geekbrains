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
	Key   string   `json:"key"`
	Links []string `json:"links"`
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

	switch req.Method {
	case "POST":
		// Обработка ниже.
	case "GET":
		res.WriteHeader(200)
		res.Write([]byte("Используй POST-запрос!"))
		return
	default:
		res.WriteHeader(404)
		res.Write([]byte(""))
		return
	}

	var err error
	var body []byte
	var resBody []byte
	var reqBody = RequesBody{}
	var okLinks []string
	var statusCode int = 200

	// Получим тело запроса.
	if body, err = ioutil.ReadAll(req.Body); err != nil {
		//fmt.Println("Не удалось прочесть тело запроса:", err)
		statusCode = 500
		resBody = []byte("Не удалось прочесть тело запроса:" + err.Error())
	} else if len(body) == 0 {
		// fmt.Println("Тело запроса пусто")
		resBody = []byte("Тело запроса пусто")
	} else if err := json.Unmarshal(body, &reqBody); err != nil {
		//fmt.Println("Не удалось разобрать JSON:", err)
		statusCode = 500
		resBody = []byte("Не удалось разобрать JSON: " + err.Error())
	} else if okLinks, err = findKey(reqBody.Key, reqBody.Links); err != nil {
		//fmt.Println(err)
		resBody = []byte(err.Error())
	} else if resBody, err = json.Marshal(okLinks); err != nil {
		//fmt.Println("Не удалось создать тело ответа:", err)
		statusCode = 500
		resBody = []byte("Не удалось создать тело ответа: " + err.Error())
	}

	res.WriteHeader(statusCode)
	res.Write(resBody)
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

			if res, err = http.Get(link); err != nil {
				err = nil
				return
			}
			defer res.Body.Close()

			if body, err = ioutil.ReadAll(res.Body); err != nil {
				err = nil
				return
			}

			if strings.Contains(string(body), key) {
				results = append(results, link)
			}
		}(links[idx])
	}

	wgr.Wait()
	return
}
