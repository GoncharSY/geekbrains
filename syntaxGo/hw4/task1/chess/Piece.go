package chess

import "fmt"

// Piece - описывает любую шахматную фигуру.
type Piece struct {
	name string
}

// NewPiece - создаст новый объект Piece и вернет указатель на него.
func NewPiece(name string) (p *Piece, err error) {
	p = new(Piece)

	if err = p.setName(name); err != nil {
		p = nil
		return
	}

	return
}

// GetName - возвращает имя фигуры.
func (p *Piece) GetName() string {
	return p.name
}

// SetName - устанавливает имя для фигуры.
// Имя может быть установлено только один раз.
// При попытке повторного присвоения имени фигуре вернется ошибка!
func (p *Piece) setName(name string) error {
	if p.name != "" {
		return fmt.Errorf("У фигуры уже есть имя: %s", p.GetName())
	}

	if name == "" {
		return fmt.Errorf("Имя фигуры не может быть пустой строкой")
	}

	p.name = name

	return nil
}
