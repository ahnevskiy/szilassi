package main

import (
	"fmt"
	"image"
	"math/rand"

	"image/color"
)

// Point is a type for point desription
type Point struct {
	X, Y, Z int
}

func generatePoint() Point {
	return Point{
		rand.Intn(spaceResolution),
		rand.Intn(spaceResolution),
		rand.Intn(1)}
	// rand.Intn(spaceResolution)}
}

func swapPoints(p1, p2 *Point) {
	buf := *p1
	*p1 = *p2
	*p2 = buf
}

// Draw is for drawing a point on canvas
func (p *Point) Draw(canvas *image.RGBA, color color.RGBA) {
	canvas.Set(p.X, p.Y, color)
}

// Mark is for marking a point with cross on canvas
func (p *Point) Mark(canvas *image.RGBA, color color.RGBA) {
	canvas.Set(p.X+1, p.Y, color)
	canvas.Set(p.X-1, p.Y, color)
	canvas.Set(p.X, p.Y+1, color)
	canvas.Set(p.X, p.Y-1, color)
}

// CheckCompare is for eqals to p and p2
func (p *Point) CheckCompare(p2 Point) bool {
	return p.X == p2.X && p.Y == p2.Y && p.Z == p2.Z
}

func (p *Point) swapCoordinates() {
	buf := p.X
	p.X = p.Y
	p.Y = buf
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d, %d)", p.X, p.Y, p.Z)
}
