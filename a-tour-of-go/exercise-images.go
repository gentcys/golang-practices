package main

import (
	"image"
	"golang.org/x/tour/pic"
	"image/color"
)

type Image struct {
	Width, Height int
	color uint8
}

func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.color + uint8(x), r.color+uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 200}
	pic.ShowImage(&m)
}
