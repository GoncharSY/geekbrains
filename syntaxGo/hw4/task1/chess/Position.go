package chess

import (
	"fmt"
)

// PositionXMin - минимальная координата по X на шахматной доске.
const PositionXMin int = 1

// PositionXMax - максимальная координата по X на шахматной доске.
const PositionXMax int = 8

// PositionYMin - минимальная координата по Y на шахматной доске.
const PositionYMin int = 1

// PositionYMax - максимальная координата по Y на шахматной доске.
const PositionYMax int = 8

// Position - описывает позицию на шахматной доске.
type Position struct {
	x int
	y int
}

// NewPosition - создаст новый объект позиции и вернет указатель на него.
func NewPosition(x, y int) (p *Position, err error) {
	p = new(Position)

	if err = p.setX(x); err != nil {
		p = nil
		return
	}

	if err = p.setY(y); err != nil {
		p = nil
		return
	}

	return
}

// GetName - вернет имя для позиции.
// Имя позиции формируется на основе ее координат на игральном поле (шахматной доске).
func (p *Position) GetName() string {
	var yName = [9]string{"", "A", "B", "C", "D", "E", "F", "G", "H"}

	return fmt.Sprintf("%s%d", yName[p.GetY()], p.GetX())
}

// GetX - вернет координату по X.
func (p *Position) GetX() int {
	return p.x
}

// GetY - вернет координату по Y.
func (p *Position) GetY() int {
	return p.y
}

// ================================================================================================
// PRIVATE PART
// ================================================================================================

// setX - установит координату по X.
func (p *Position) setX(x int) error {
	if x < PositionXMin || x > PositionXMax {
		return fmt.Errorf("Выход за допустимый диапазон (от %d до %d)", PositionXMin, PositionXMax)
	}

	p.x = x

	return nil
}

// setY - установит координату по Y.
func (p *Position) setY(y int) error {
	if y < PositionYMin || y > PositionYMax {
		return fmt.Errorf("Выход за допустимый диапазон (от %d до %d)", PositionYMin, PositionYMax)
	}

	p.y = y

	return nil
}
