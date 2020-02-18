package main

import (
	"fmt"

	"./chess"
)

// Описывает методы именованных сущностей в шахматах.
type chessEntity interface {
	GetName() string
}

func main() {
	var chessPiece *chess.Piece
	var chessPositions = make([]*chess.Position, 2)
	var chessEntities []chessEntity

	chessPiece, _ = chess.NewPiece("Пешка")
	chessPositions[0], _ = chess.NewPosition(2, 5)
	chessPositions[1], _ = chess.NewPosition(4, 5)

	chessEntities = []chessEntity{
		chessPiece,
		chessPositions[0],
		chessPositions[1],
		chess.NewMove(chessPositions[0], chessPositions[1]),
	}

	fmt.Println("Именованные сущности в шахматах:")
	printNames(chessEntities...)
}

// Вывод имен сущностей.
func printNames(items ...chessEntity) {
	for i := range items {
		switch items[i].(type) {
		case *chess.Piece:
			fmt.Print(" Фигура: ")
		case *chess.Position:
			fmt.Print("Позиция: ")
		case *chess.Move:
			fmt.Print("    Ход: ")
		}

		fmt.Println(items[i].GetName())
	}
}
