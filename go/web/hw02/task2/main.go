package main

import (
	"fmt"
	"log"
	"net/http"
)

var myCookie = http.Cookie{
	Name:   "Name_of_person",
	Value:  "No name",
	MaxAge: 300,
}

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
	return
}

// Обработать запрос к корневому маршруту.
func onRequestRoot(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	if req.Method != "GET" {
		res.WriteHeader(404)
		res.Write([]byte(""))
		return
	}

	res.WriteHeader(200)
	res.Write([]byte("Выберите действие: get/set/delete cookie"))
}

// Обработать запрос на установку cookie.
func onSetCookie(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	if req.Method != "GET" {
		res.WriteHeader(404)
		res.Write([]byte(""))
		return
	}

	if cookieValue := req.FormValue("value"); len(cookieValue) != 0 {
		myCookie.Value = cookieValue
	}

	myCookie.MaxAge = 300
	http.SetCookie(res, &myCookie)

	res.WriteHeader(200)
	res.Write([]byte("Cookie '" + myCookie.Name + "' установлен, используйте /get, чтобы получить его значение"))
}

// Обработать запрос на получение cookie.
func onGetCookie(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	var err error
	var cookie *http.Cookie

	if req.Method != "GET" {
		res.WriteHeader(404)
		res.Write([]byte(""))
		return
	}

	if cookie, err = req.Cookie(myCookie.Name); err != nil {
		res.WriteHeader(200)
		res.Write([]byte("Cookie не найден: " + err.Error()))
		return
	}

	res.WriteHeader(200)
	res.Write([]byte("Найдены cookie с именем '" + cookie.Name + "' и значением '" + cookie.Value + "'"))
}

// Обработать запрос на удаление cookie.
func onDeleteCookie(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	if req.Method != "GET" {
		res.WriteHeader(404)
		res.Write([]byte(""))
		return
	}

	myCookie.MaxAge = -1
	http.SetCookie(res, &myCookie)

	res.WriteHeader(200)
	res.Write([]byte("Cookie c именем '" + myCookie.Name + "' был удален"))
}
