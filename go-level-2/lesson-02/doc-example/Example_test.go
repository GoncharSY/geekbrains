package docexample_test

import docexample "go-level-2/lesson-02/doc-example"

// Пример использования инструментов пакета.
// Этот пример создан специально для проверки его отображения в документации
// создаваемой через утилиту `godoc`.
func Example() {
	var obj1 = docexample.MyStruct{
		Id:   "1",
		Name: "Object-1",
		Date: "01.01.2022",
	}

	obj1.MyMethod() // true
}
