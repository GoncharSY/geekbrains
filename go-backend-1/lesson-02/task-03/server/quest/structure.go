package quest

import (
	"fmt"
	"math/rand"
	"time"
)

// Структура игровой задачи.
type Structure struct {
	Operator Operator
	Operand1 int
	Operand2 int
}

// Преобразовать задачу в текст.
func (qst *Structure) String() string {
	return fmt.Sprintf("%v %v %v = ?", qst.Operand1, qst.Operator, qst.Operand2)
}

// Сбросить текущее условие задачи и создать новое.
func (qst *Structure) Reset() {
	var src = rand.NewSource(time.Now().UnixNano())
	var gen = rand.New(src)

	qst.Operator = Operator(gen.Intn(int(Div)) + 1)
	qst.Operand1 = gen.Intn(100)
	qst.Operand2 = gen.Intn(100)

	for qst.Operator == Div && qst.Operand2 == 0 {
		qst.Operand2 = rand.Int()
	}
}

// Проверить, является ли указанное значение решением задачи.
func (qst *Structure) IsSolution(val int) bool {
	return val == qst.answer()
}

// Получить ответ текущей игровой задачи.
func (qst *Structure) answer() int {
	switch qst.Operator {
	case Inc:
		return qst.Operand1 + qst.Operand2
	case Dec:
		return qst.Operand1 - qst.Operand2
	case Mul:
		return qst.Operand1 * qst.Operand2
	case Div:
		return qst.Operand1 / qst.Operand2
	default:
		panic("unknown operand")
	}
}
