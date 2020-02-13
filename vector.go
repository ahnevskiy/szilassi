package main

import (
	"math"
)

// Vector is a type for desript vector
type Vector struct {
	coordinates Point
}

// Length returns length of vector
func (v *Vector) Length() float64 {
	x, y, z := float64(v.coordinates.X), float64(v.coordinates.Y), float64(v.coordinates.Z)
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2) + math.Pow(z, 2))
}

func angleBetweenVectors(v1, v2 Vector) float64 {
	return math.Acos(DotProduct(v1, v2) / (v1.Length() * v2.Length()))
}

// CrossProduct is calc a cross product of vectors
func CrossProduct(v1, v2 Vector) Vector {
	resultVectorCoordinates := Point{
		v1.coordinates.Y*v2.coordinates.Z - v1.coordinates.Z*v2.coordinates.Y,
		v1.coordinates.Z*v2.coordinates.X - v1.coordinates.X*v2.coordinates.Z,
		v1.coordinates.X*v2.coordinates.Y - v1.coordinates.Y*v2.coordinates.X}
	return Vector{resultVectorCoordinates}
}

// DotProduct is calc a dot product of vectors
func DotProduct(v1, v2 Vector) float64 {
	x1, y1, z1 := float64(v1.coordinates.X), float64(v1.coordinates.Y), float64(v1.coordinates.Z)
	x2, y2, z2 := float64(v2.coordinates.X), float64(v2.coordinates.Y), float64(v2.coordinates.Z)
	return x1*x2 + y1*y2 + z1*z2
}

// PseudoDotProduct is calc pseudo* dot product of vectors
// * https://ru.wikipedia.org/wiki/Псевдоскалярное_произведение
func PseudoDotProduct(v1, v2 Vector) float64 {
	x1, y1, z1 := float64(v1.coordinates.X), float64(v1.coordinates.Y), float64(v1.coordinates.Z)
	x2, y2, z2 := float64(v2.coordinates.X), float64(v2.coordinates.Y), float64(v2.coordinates.Z)
	z1 = z2
	z2 = z1
	return x1*y2 - x2*y1
}
