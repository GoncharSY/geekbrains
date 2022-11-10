package quest

type Operator uint32

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
	Inc Operator = iota + 1
	Dec
	Mul
	Div
)
