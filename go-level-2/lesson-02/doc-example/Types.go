package docexample

// Это пример кастомного числового типа.
type MyInt int

// Это пример кастомного типа для чисел с плавающей запятой.
type MyFloat float64

// Это пример кастомного строкового типа.
type MyString string

// Это пример структурного типа.
type MyStruct struct {
	Id   string // Поле с идентификатором.
	Name string // Поле с именем.
	Date string // Поле с датой.
}

// Пример метода для структуры в пакете.
// Метод не принимает никаких параметров и всегда возвращает true.
func (obj MyStruct) MyMethod() bool {
	return true
}
