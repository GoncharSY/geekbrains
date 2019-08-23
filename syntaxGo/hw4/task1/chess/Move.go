package chess

import "fmt"

// Move - описывает ход в игре.
type Move struct {
	from *Position
	to   *Position
}

// NewMove - создаст новый ход и вернет указатель на объект.
func NewMove(from, to *Position) (m *Move) {
	m = new(Move)
	m.from = from
	m.to = to

	return
}

// GetName - венет имя игрового хода.
// Имя сформируется на основе исходной и конечной позиций.
func (m *Move) GetName() string {
	return fmt.Sprintf("%s - %s", m.from.GetName(), m.to.GetName())
}
