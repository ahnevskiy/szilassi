package main

import (
	"flag"
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"
)

var (
	// Variables for input arguments
	spaceResolution, vertexesCount int
	outputFileName                 string

	// Any global variables
	precision int

	// Constants for basic colors
	colorWhite = color.RGBA{255, 255, 255, 255}
	colorBlack = color.RGBA{0, 0, 0, 255}
	colorRed   = color.RGBA{255, 0, 0, 255}
	colorGreen = color.RGBA{0, 255, 0, 255}
	colorBlue  = color.RGBA{0, 0, 255, 255}
	colorTeal  = color.RGBA{0, 200, 200, 255}
)

const ()

func argsParse() {
	// Parse arguments
	flag.StringVar(&outputFileName, "img", "map.png", "The name of output image with result")
	flag.IntVar(&spaceResolution, "r", 1000, "The resolution of space")
	flag.IntVar(&vertexesCount, "v", 6, "The count of vertexes in polygon")
	flag.Parse()

	// Print arguments in terminal
	fmt.Printf("Resolution: [%d]\n", spaceResolution)
	fmt.Printf("Vertexes:   [%d]\n", vertexesCount)
	fmt.Printf("Output:     [%s]\n", outputFileName)
}

func init() {
	rand.Seed(time.Now().Unix())
	argsParse()
	precision = int(math.Log10(float64(spaceResolution))) - 1
}

func main() {
	// Create an image for output
	img := Canvas{}
	img.Create(outputFileName, colorBlack)
	defer img.SavePNG()

	// Generate and draw a correct polygon
	pol := generatePolygon(vertexesCount)
	pol.Draw(img.canvas, colorBlue)

}
