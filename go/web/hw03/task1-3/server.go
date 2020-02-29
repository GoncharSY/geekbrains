package main

import (
	"fmt"
	"log"
	"net/http"

	"./db"
)

func main() {
	var port string = "8080"
	var srvMux *http.ServeMux

	srvMux = http.NewServeMux()
	srvMux.HandleFunc("/blog", onRequestBlog)
	srvMux.HandleFunc("/post/{idx}", onRequestPost)
	srvMux.HandleFunc("/", onRequestRoot)

	fmt.Println("Запуск сервера...")
	log.Fatal(http.ListenAndServe(":"+port, srvMux))
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

	var err error
	var blog *db.Blog

	if blog, err = db.GetBlog(); err != nil {
		res.WriteHeader(404)
		res.Write([]byte("Блог не доступен в настоящее время"))
		return
	}

	res.Write([]byte("Обращение к блогу '" + blog.Name + "'"))
}

// Обработать запрос для обращения к посту в блоге.
func onRequestPost(res http.ResponseWriter, req *http.Request) {
	onRequest(res, req)

	var err error
	var blog *db.Blog
	var post *db.Post
	var postIdx = 0
	var postCnt = 0

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
	res.Write([]byte(fmt.Sprintf("Обращение к посту #%v от автора %v", postIdx, post.Author)))
}
