package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	var fs = http.FileServer(http.Dir("."))

	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Отправляет страницу с приветствием, созданную на основе шаблона.
func hello(res http.ResponseWriter, req *http.Request) {
	var query = req.URL.Query()
	var tmplPath = "hello.html"
	var tmpl *template.Template
	var err error
	var name string

	// Получим шаблон
	tmpl, err = template.ParseFiles(tmplPath)
	if err != nil {
		log.Print("ERROR:", err)
		res.WriteHeader(500)
		res.Write([]byte(err.Error()))
		return
	}

	// Получим имя
	name = query.Get("name")
	if name == "" {
		name = "NoName"
	}

	// Установим заголовки
	res.Header().Set("Content-Type", "text/html")

	// Отправим ответ на запрос.
	tmpl.Execute(res, name)
}
