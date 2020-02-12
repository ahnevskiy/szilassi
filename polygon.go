package main

import (
	"image"
	"image/color"
)

// Polygon is a type for polygon desription
type Polygon struct {
	vertexes [vertexesCount]Point
	edges    [vertexesCount]Segment
}

func generatePolygon() Polygon {
	polygon := Polygon{}
	polygon.vertexes[0] = generatePoint()
	newVertex := Point{}
	for vertexNumber := 1; vertexNumber < vertexesCount; vertexNumber++ {
		newVertex = generatePoint()
		isVertexOnPrevEdge := false
		if vertexNumber > 1 {
			isVertexOnPrevEdge = polygon.edges[vertexNumber-1].IsPointBelongLine(newVertex)
		}
		for polygon.vertexes[vertexNumber-1].Compare(newVertex) == true || isVertexOnPrevEdge {
			newVertex = generatePoint()
			if vertexNumber > 1 {
				isVertexOnPrevEdge = polygon.edges[vertexNumber-1].IsPointBelongLine(newVertex)
			}
		}
		polygon.vertexes[vertexNumber] = newVertex
		polygon.edges[vertexNumber-1] = Segment{polygon.vertexes[vertexNumber-1],
			polygon.vertexes[vertexNumber]}
	}
	polygon.edges[vertexesCount-1] = Segment{polygon.vertexes[vertexesCount-1], polygon.vertexes[0]}
	return polygon
}

// Draw is for drawing polygon on canvas
func (p *Polygon) Draw(canvas *image.RGBA, color color.RGBA) {
	for i := 0; i < vertexesCount; i++ {
		// fmt.Printf("e%d v1=%s \n", i, p.vertexes[i].String())
		p.edges[i].Draw(canvas, color)
	}
}
