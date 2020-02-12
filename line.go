package main

import (
	"image"
	"image/color"
	"math"
)

// Segment is a type for desript segment/line with a two points
type Segment struct {
	p1, p2 Point
}

// CrossProduct is calc a cross product of vectors
func CrossProduct(s1, s2 Segment) Point {
	v1 := s1.Vector()
	v2 := s2.Vector()
	return Point{
		v1.Y*v2.Z - v1.Z*v2.Y,
		v1.Z*v2.X - v1.X*v2.Z,
		v1.X*v2.Y - v1.Y*v2.X}
}

// IsPointBelongLine return True if point belongs to line
func (s *Segment) IsPointBelongLine(p Point) bool {
	deltaX := s.p2.X - s.p1.X
	deltaY := s.p2.Y - s.p1.Y
	deltaZ := s.p2.Z - s.p1.Z
	xPart := precisionRound(float64(p.X-s.p1.X)/float64(deltaX), precision)
	yPart := precisionRound(float64(p.Y-s.p1.Y)/float64(deltaY), precision)
	zPart := precisionRound(float64(p.Z-s.p1.Z)/float64(deltaZ), precision)
	if deltaX == 0 {
		if deltaY == 0 {
			return p.X == s.p1.X && p.Y == s.p1.Y
		}
		if deltaZ == 0 {
			return p.X == s.p1.X && p.Z == s.p1.Z
		}
		return yPart == zPart
	}
	if deltaY == 0 {
		if deltaX == 0 {
			return p.Y == s.p1.Y && p.X == s.p1.X
		}
		if deltaZ == 0 {
			return p.Y == s.p1.Y && p.Z == s.p1.Z
		}
		return xPart == zPart
	}
	if deltaZ == 0 {
		if deltaX == 0 {
			return p.Z == s.p1.Z && p.X == s.p1.X
		}
		if deltaY == 0 {
			return p.Z == s.p1.Z && p.Y == s.p1.Y
		}
		return xPart == yPart
	}
	return (xPart == yPart) && (xPart == zPart)
}

// Draw is for drawing segment on canvas
// Bresenham's line algorithm
func (s *Segment) Draw(canvas *image.RGBA, color color.RGBA) {

	p1 := s.p1
	p2 := s.p2

	mirror := false
	deltaX := int(math.Abs(float64(p2.X - p1.X)))
	deltaY := int(math.Abs(float64(p2.Y - p1.Y)))

	if deltaX < deltaY {
		p1.swapCoordinates()
		p2.swapCoordinates()
		deltaX = int(math.Abs(float64(p2.X - p1.X)))
		deltaY = int(math.Abs(float64(p2.Y - p1.Y)))
		mirror = true
	}

	if p1.X > p2.X {
		swapPoints(&p1, &p2)
	}

	dirX := p2.X - p1.X
	if dirX < 0 {
		dirX = -1
	} else {
		dirX = 1
	}

	dirY := p2.Y - p1.Y
	if dirY < 0 {
		dirY = -1
	} else {
		dirY = 1
	}

	err := 0
	deltaErr := deltaY + 1
	cursor := Point{p1.X, p1.Y, p1.Z}
	for i := 0; cursor.X <= p2.X; cursor.X += dirX {
		i++
		if !mirror {
			cursor.Draw(canvas, color)
		} else {
			mirroredCursor := Point{cursor.Y, cursor.X, cursor.Z}
			mirroredCursor.Draw(canvas, color)
		}
		err += deltaErr
		if err >= (deltaX + 1) {
			cursor.Y += dirY
			err -= deltaX + 1
		}
	}
}

// Vector return coordinates of vector of line "s"
func (s *Segment) Vector() Point {
	return Point{
		s.p1.X - s.p2.X,
		s.p1.Y - s.p2.Y,
		s.p1.Z - s.p2.Z}
}
