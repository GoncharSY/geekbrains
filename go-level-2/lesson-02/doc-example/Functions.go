package docexample

/*
Пример функции, которая возвращает строку.
Эта функция принимает три параметра:
	- num - какое-то целое число
	- str - какая-то строка
	- flag - флажок (true или false)
*/
func MyFunction(num int, str string, flag bool) string {
	if flag && num > 0 {
		return str
	}

	return ""
}
