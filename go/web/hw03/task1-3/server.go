package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"./db"
)

func main() {
	var port string = "8080"
	var router *mux.Router

	router = mux.NewRouter()
	router.HandleFunc("/blog", onRequestBlog)
	router.HandleFunc("/post/{idx:[0-9]+}", onRequestPost)
	router.HandleFunc("/", onRequestRoot)

	fmt.Println("Запуск сервера...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Выполнить предварительную обработку запроса.
func onRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Пришел запрос: %s %s\n", req.Method, req.URL.Path)
	res.Header().Set("Access-Control-Allow-Origin", "*")
}

// Обработать запрос по корневому маршруту.
func onRequestRoot(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	res.Write([]byte("Обращение к корневому маршруту"))
}

// Обработать запрос для обращения к блогу.
func onRequestBlog(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	if req.Method != "GET" {
		res.WriteHeader(404)
		res.Write([]byte("Доступен только метод запроса GET"))
		return
	}

	var err error
	var blog *db.Blog

	if blog, err = db.GetBlog(); err != nil {
		fmt.Println(err)
		res.WriteHeader(404)
		res.Write([]byte("Блог не доступен в настоящее время"))
		return
	}

	var tmp *template.Template
	var tmpFile = ".\\tmp\\blog"

	if tmp, err = tmp.ParseFiles(tmpFile); err != nil {
		res.WriteHeader(500)
		res.Write([]byte("Ошибка шаблона: " + err.Error()))
		return
	}

	tmp.ExecuteTemplate(res, "Blog", blog)
}

// Обработать запрос для обращения к посту в блоге.
func onRequestPost(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	var err error
	var blog *db.Blog
	var post *db.Post

	var varsMap map[string]string
	var postIdx = 0
	var postCnt = 0

	varsMap = mux.Vars(req)
	if idx, ok := varsMap["idx"]; !ok {
		res.WriteHeader(400)
		res.Write([]byte("Пост не найден"))
		return
	} else if postIdx, err = strconv.Atoi(idx); err != nil {
		res.WriteHeader(400)
		res.Write([]byte("Пост не найден: " + err.Error()))
		return
	}

	if blog, err = db.GetBlog(); err != nil {
		res.WriteHeader(404)
		res.Write([]byte("Блог не доступен в настоящее время"))
		return
	}

	postCnt = len(blog.Posts)
	if postCnt == 0 {
		res.WriteHeader(404)
		res.Write([]byte("В данном блоге нет записей"))
		return
	} else if postIdx < 0 || postIdx >= postCnt {
		res.WriteHeader(404)
		res.Write([]byte("Пост не найден"))
		return
	}

	post = &blog.Posts[postIdx]
	res.Write([]byte(fmt.Sprintf("Обращение к посту #%v с названием '%v'", postIdx, post.Name)))
}
