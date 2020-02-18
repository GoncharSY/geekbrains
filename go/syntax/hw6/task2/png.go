package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	var size, indent = 500, 50
	var white = color.RGBA{255, 255, 255, 255}
	var black = color.RGBA{0, 0, 0, 255}
	var img = image.NewRGBA(image.Rect(0, 0, size, size))
	var file, err = os.Create("rectangle.png")

	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()

	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)

	for step := 1; (step * indent) <= (size - indent); step++ {
		stepInd := step * indent

		for pix := indent; pix < (size - indent); pix++ {
			img.Set(pix, stepInd, black) // horizontal
			img.Set(stepInd, pix, black) // vertical
		}
	}

	png.Encode(file, img)
}
