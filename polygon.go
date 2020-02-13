package main

import (
	"image"
	"image/color"
)

// Polygon is a type for polygon desription
type Polygon struct {
	vertexes      []Point
	edges         []Segment
	vertexesCount int
}

func generateVertex(pol *Polygon, vertexNumber int) (Point, bool) {
	newVertex := generatePoint()
	newEdge := Segment{pol.vertexes[vertexNumber-1], newVertex}
	isVertexOnPrevEdge := false
	isVertexOnOtherVertexes := false
	isEdgeIntersectWithOtherEdges := false
	if vertexNumber > 1 {
		isVertexOnPrevEdge = pol.edges[vertexNumber-1].IsPointBelongLine(newVertex)
	}
	for j := 0; j < vertexNumber && !isVertexOnOtherVertexes; j++ {
		isVertexOnOtherVertexes = pol.vertexes[j].CheckCompare(newVertex)
	}
	for k := 0; k < vertexNumber && !isVertexOnOtherVertexes; k++ {
		isVertexOnOtherVertexes = pol.edges[k].CheckIntersection(newEdge)
	}
	return newVertex, isVertexOnPrevEdge || isVertexOnOtherVertexes || isEdgeIntersectWithOtherEdges
}

func generatePolygon(vertexesCount int) Polygon {
	polygon := Polygon{}
	polygon.vertexesCount = vertexesCount
	isLastEdgeIntersectPolygon := true
	for isLastEdgeIntersectPolygon {
		polygon.vertexes = make([]Point, polygon.vertexesCount)
		polygon.edges = make([]Segment, polygon.vertexesCount)

		polygon.vertexes[0] = generatePoint()
		for vertexNumber := 1; vertexNumber < polygon.vertexesCount; vertexNumber++ {
			newVertex, isSituableVertex := generateVertex(&polygon, vertexNumber)
			for isSituableVertex {
				newVertex, isSituableVertex = generateVertex(&polygon, vertexNumber)
			}
			polygon.vertexes[vertexNumber] = newVertex
			polygon.edges[vertexNumber-1] = Segment{polygon.vertexes[vertexNumber-1],
				polygon.vertexes[vertexNumber]}
		}
		polygon.edges[polygon.vertexesCount-1] = Segment{polygon.vertexes[polygon.vertexesCount-1], polygon.vertexes[0]}
		isLastEdgeIntersectPolygon = false
		for i := 0; i < vertexesCount-1 && !isLastEdgeIntersectPolygon; i++ {
			isLastEdgeIntersectPolygon = polygon.edges[i].CheckIntersection(polygon.edges[polygon.vertexesCount-1])
		}
	}
	return polygon
}

// Draw is for drawing polygon on canvas
func (p *Polygon) Draw(canvas *image.RGBA, color color.RGBA) {
	for i := 0; i < vertexesCount; i++ {
		p.edges[i].Draw(canvas, color)
	}
	for i := 0; i < vertexesCount; i++ {
		p.vertexes[i].Mark(canvas, colorRed)
	}
}
