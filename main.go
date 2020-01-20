package main

import (
	"fmt"
	"time"
)

type _Point struct {
	X, Y, Z float64
}

func reprPoint(p _Point) {
	fmt.Printf("X: %.5f\n", p.X)
	fmt.Printf("Y: %.5f\n", p.Y)
	fmt.Printf("Z: %.5f\n", p.Z)
}
func plus(p, q _Point, r *_Point) {
	r.X = p.X + q.X
	r.Y = p.Y + q.Y
	r.Z = p.Z + q.Z
}

func init() {

}

func main() {
	p1 := _Point{1, 2, 3}
	p2 := _Point{2, 3, 4}
	c := _Point{0, 0, 0}
	d := _Point{0, 0, 0}
	go plus(p1, p2, &c)
	go plus(p2, p1, &d)
	time.Sleep(1 * time.Second)
	reprPoint(c)
	reprPoint(d)

}
