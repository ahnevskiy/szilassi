package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"
)

var (
	precision = int(math.Log10(float64(spaceResolution))) - 1

	colorWhite = color.RGBA{255, 255, 255, 255}
	colorBlack = color.RGBA{0, 0, 0, 255}
	colorRed   = color.RGBA{255, 0, 0, 255}
	colorGreen = color.RGBA{0, 255, 0, 255}
	colorBlue  = color.RGBA{0, 0, 255, 255}
	colorTeal  = color.RGBA{0, 200, 200, 255}
)

const (
	spaceResolution = 1000
	vertexesCount   = 5
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	img := Canvas{}
	img.Create("map.png")
	defer img.SavePNG()

	pol := generatePolygon()
	pol.Draw(img.canvas, colorBlue)
}
