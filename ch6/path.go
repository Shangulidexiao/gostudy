package ch6

import (
	"math"
)

type Point struct {
	X, Y float64
}

type Path struct {
	P1, P2 Point
}

func PointDistance(p1, p2 Point) float64 {
	return math.Hypot(p1.X-p2.X, p1.Y-p2.Y)
}

func (p Point) Distance(p1 Point) float64 {
	return math.Hypot(p.X-p1.X, p.Y-p1.Y)
}

func (p *Path) Distance() float64 {
	return math.Hypot(p.P1.X-p.P2.X, p.P1.Y-p.P2.Y)
}

func (p *Path) ChangeX(x float64) {
	p.P1.X = x
}

func (p *Point) ChangeX(x float64) {
	p.X = x
}
