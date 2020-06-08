package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var dataPath = ".\\db\\data.json"
var blog = Blog{}
var errDB = fmt.Errorf("База данных не инициализинована")

// Инициализировать данные.
func init() {
	var err error
	var data []byte

	if data, err = ioutil.ReadFile(dataPath); err != nil {
		fmt.Println("Ошибка чтения файла данных БД: " + err.Error())
		return
	}

	if err = json.Unmarshal(data, &blog); err != nil {
		fmt.Println("Ошибка распаковки данных БД: " + err.Error())
		return
	}

	errDB = nil
	fmt.Println("База данных блога инициализирована.")
}

// GetBlog - получить блог.
func GetBlog() (*Blog, error) {
	return &blog, errDB
}

// Blog - описывает блог со списком постов.
type Blog struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Posts  []Post `json:"posts"`
}

// Post - описывает отдельный пост в блоге.
type Post struct {
	Name      string `json:"name"`
	Author    string `json:"author"`
	CreatedAt string `json:"createdAt"`
	Text      string `json:"text"`
	Comments  []Post `json:"comments"`
}

// GetCommentsCount - возвращает количество комментариев к посту.
func (p *Post) GetCommentsCount() int {
	return len(p.Comments)
}
