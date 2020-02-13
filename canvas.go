package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// Canvas is a type for painting image
type Canvas struct {
	canvas *image.RGBA
	file   *os.File
}

// Create is a method for creating canvas for PNG image
func (c *Canvas) Create(fileName string, fontColor color.RGBA) {
	c.file, _ = os.Create(fmt.Sprintf("%s", fileName))
	// if err != nil {
	// 	fmt.Errorf("%s", err)
	// }
	c.canvas = image.NewRGBA(image.Rect(0, 0, spaceResolution, spaceResolution))
	draw.Draw(c.canvas, c.canvas.Bounds(), &image.Uniform{fontColor}, image.ZP, draw.Src)
}

// SavePNG is a method for saving result canvas in file
func (c *Canvas) SavePNG() {
	png.Encode(c.file, c.canvas)
}
