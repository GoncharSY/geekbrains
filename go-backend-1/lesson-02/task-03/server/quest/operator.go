package quest

// Оператор для игровой задачи.
type Operator uint32

// Преобразовать оператор в строку.
func (opr Operator) String() string {
	switch opr {
	case Div:
		return "/"
	case Mul:
		return "*"
	case Dec:
		return "-"
	default:
		return "+"
	}
}

const (
	Inc Operator = iota + 1 // Оператор сложения.
	Dec                     // Оператор вычитания.
	Mul                     // Оператор умножения.
	Div                     // Оператор деления.
)
