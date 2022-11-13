package quest

// Создать новую игровую задачу.
func New() *Structure {
	qst := &Structure{}
	qst.Reset()
	return qst
}
