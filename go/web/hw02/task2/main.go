package main

import (
	"fmt"
	"log"
	"net/http"
)

var nameOfCookie = "Name_of_person"

func main() {
	var srvMux *http.ServeMux

	srvMux = http.NewServeMux()
	srvMux.HandleFunc("/", onRequestRoot)
	srvMux.HandleFunc("/set", onSetCookie)
	srvMux.HandleFunc("/get", onGetCookie)
	srvMux.HandleFunc("/delete", onDeleteCookie)

	fmt.Println("Запуск сервера...")
	log.Fatal(http.ListenAndServe(":8080", srvMux))
}

// Предварительная обработка запроса.
func onRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Пришел запрос:", req.URL.Path)
	fmt.Println("Метод запроса:", req.Method)

	res.Header().Set("Access-Control-Allow-Origin", "*")
}

// Обработать запрос к корневому маршруту.
func onRequestRoot(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	if req.Method != "GET" {
		res.WriteHeader(404)
		res.Write([]byte("Доступен только метод запроса GET"))
		return
	}

	res.Write([]byte("Выберите действие: get/set/delete cookie"))
}

// Обработать запрос на установку cookie.
func onSetCookie(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	if req.Method != "GET" {
		res.WriteHeader(404)
		res.Write([]byte("Доступен только метод запроса GET"))
		return
	}

	var cookie = http.Cookie{
		Name:   nameOfCookie,
		Value:  "No name",
		MaxAge: 300,
	}

	if valueOfCookie := req.FormValue("value"); len(valueOfCookie) != 0 {
		cookie.Value = valueOfCookie
	}

	cookie.MaxAge = 300
	http.SetCookie(res, &cookie)
	res.Write([]byte("Cookie '" + cookie.Name + "' установлен, используйте /get, чтобы получить его значение"))
}

// Обработать запрос на получение cookie.
func onGetCookie(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	if req.Method != "GET" {
		res.WriteHeader(404)
		res.Write([]byte("Доступен только метод запроса GET"))
		return
	}

	var err error
	var cookie *http.Cookie

	if cookie, err = req.Cookie(nameOfCookie); err != nil {
		res.Write([]byte("Cookie не найден: " + err.Error()))
		return
	}

	res.Write([]byte("Найден cookie с именем '" + cookie.Name + "' и значением '" + cookie.Value + "'"))
}

// Обработать запрос на удаление cookie.
func onDeleteCookie(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	if req.Method != "GET" {
		res.WriteHeader(404)
		res.Write([]byte("Доступен только метод запроса GET"))
		return
	}

	var err error
	var cookie *http.Cookie

	if cookie, err = req.Cookie(nameOfCookie); err != nil {
		res.Write([]byte("Cookie не найден: " + err.Error()))
		return
	}

	cookie.MaxAge = -1
	http.SetCookie(res, cookie)
	res.Write([]byte("Cookie c именем '" + cookie.Name + "' был удален"))
}
