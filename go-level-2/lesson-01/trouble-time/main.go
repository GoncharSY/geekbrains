// Это код программы для первого задания.
// - Паническая ситуация создается.
// - Паника перехватывается.
// - Создается собственная структура ошибки.
// - Время возникновения паники фиксируется.
// - В 'main' паника попадает, как обычная ошибка и там обрабатывается.

package main

import (
	"fmt"
	"time"
)

func main() {
	if err := makeImpPanic(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("run complete")
	}
}

// Создает неявную панику, обращаясь к несущствующему элементу среза.
// Перед заверешением отлавливает паническую ситуацию и инициализирует ошибку.
func makeImpPanic() (err error) {
	var slc = make([]int, 1)

	defer recoverImpPanic(&err)
	slc[1] = 7
	return err
}

// Отлавливает паническую ситуацию и превращает ее в обычную ошибку (проблему).
func recoverImpPanic(err *error) {
	if pnc := recover(); pnc != nil {
		*err = NewTrouble("panic trouble", pnc)
	}
}

// Тип для описания проблемы при выполнении программы.
// Структура проблемы реализует интерфейс 'error'.
type Trouble struct {
	Err  error
	Text string
	Time time.Time
}

// Error возвращает текст описания проблемы (ошибки).
func (t Trouble) Error() string {
	var stamp = t.Time.Format(time.StampMilli)

	if t.Err == nil {
		return fmt.Sprintf("%s %s", stamp, t.Text)
	} else {
		return fmt.Sprintf("%s %s: %s", stamp, t.Text, t.Err.Error())
	}
}

// NewTrouble создает новый объект проблемы и возвращает указатель на него.
func NewTrouble(text string, err any) *Trouble {
	return &Trouble{
		Time: time.Now(),
		Text: text,
		Err:  fmt.Errorf("%s", err),
	}
}
