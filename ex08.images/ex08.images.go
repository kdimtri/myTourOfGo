/*
Exercise: Images
Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 256, 255} in this one. */
package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

//Image interface implimentation
type Image struct {
	dx, dy int
}

//ColorModel of my Image interface
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

//Bounds of new Image interface
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.dx, img.dy)
}

//At calculates and reterns color data at specified point(x,y)
func (img Image) At(x int, y int) color.Color {
	c := uint8((x ^ y) * (x * y) / 4)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	img := Image{256, 256}
	pic.ShowImage(img)
}
