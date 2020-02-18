package queue

// StringQ описывает очередь строк.
type StringQ struct {
	items []string
}

// Add добавляет элементы в очередь.
func (q *StringQ) Add(items ...string) {
	q.items = append(q.items, items...)
}

// Take забирает элемент из очереди.
func (q *StringQ) Take() (item string, ok bool) {
	if len(q.items) == 0 {
		return
	}

	item = q.items[0]
	ok = true
	q.items = q.items[1:]

	return
}

// Len возвращает длину очереди.
func (q *StringQ) Len() int {
	return len(q.items)
}
