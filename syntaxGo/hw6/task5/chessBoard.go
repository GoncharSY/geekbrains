package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func main() {
	// размеры
	var size = 500
	var indent = 50
	var border = 5
	var flSize = size - 2*indent // игровая область (playing field size)
	var sqSize = int(flSize / 8) // квадрат (quare size)
	// цвета
	var white = color.RGBA{255, 255, 255, 255}
	var black = color.RGBA{0, 0, 0, 255}
	// элементы
	var rgbaChBoard = image.NewRGBA(image.Rect(0, 0, size, size))  // доска
	var rgbaChSq = image.NewRGBA(image.Rect(0, 0, sqSize, sqSize)) // квадрат
	var rgbaVB = image.NewRGBA(image.Rect(0, 0, border, flSize+2*border))
	var rgbaHB = image.NewRGBA(image.Rect(0, 0, flSize+2*border, border))
	// подписи
	var labels = [8]string{"A", "B", "C", "D", "E", "F", "G", "H"}

	// Раскрасим элементы
	draw.Draw(rgbaChBoard, rgbaChBoard.Bounds(), &image.Uniform{white}, image.ZP, draw.Src) // белая доска
	draw.Draw(rgbaChSq, rgbaChSq.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)       // черный квадрат
	draw.Draw(rgbaVB, rgbaVB.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)           // верт. рамка
	draw.Draw(rgbaHB, rgbaHB.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)           // гориз. рамка

	// Расставим квадраты на доску.
	for i := 0; i < 8; i++ {
		var posX = indent + i*sqSize

		for j := (i + 1) % 2; j < 8; j += 2 {
			var posY = indent + j*sqSize

			draw.Draw(rgbaChBoard, image.Rect(posX, posY, posX+sqSize, posY+sqSize), rgbaChSq, image.ZP, draw.Src)
		}
	}

	// Добавим границы игровой области.
	var left, right = indent - border, size - indent + border
	var top, bottom = indent - border, size - indent + border
	draw.Draw(rgbaChBoard, image.Rect(left, top, indent, bottom), rgbaVB, image.ZP, draw.Src)        // левая
	draw.Draw(rgbaChBoard, image.Rect(size-indent, top, right, bottom), rgbaVB, image.ZP, draw.Src)  // правая
	draw.Draw(rgbaChBoard, image.Rect(left, top, right, indent), rgbaHB, image.ZP, draw.Src)         // верхняя
	draw.Draw(rgbaChBoard, image.Rect(left, size-indent, right, bottom), rgbaHB, image.ZP, draw.Src) // нижняя

	// Добавим подписи клеткам.
	for i := 8; i > 0; i-- {
		var posX = indent / 2
		var posY = size - (indent + i*sqSize - sqSize/2) + 2*border

		addLabel(rgbaChBoard, posX, posY, strconv.Itoa(i), black)                // 1-8
		addLabel(rgbaChBoard, size-posY+2*border, size-posX, labels[i-1], black) // A-H
	}

	// Сохраним результат в файл.
	file, err := os.Create("chessBoard.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rgbaChBoard)
}

// Добавляет надпись на изображение.
func addLabel(img *image.RGBA, x, y int, label string, color color.RGBA) {
	var point = fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64),
	}

	var drawer = &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color),
		Face: basicfont.Face7x13,
		Dot:  point,
	}

	drawer.DrawString(label)
}
