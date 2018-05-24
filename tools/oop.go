package tools

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type IPoint struct{ X, Y float64 }
type Path []IPoint
type ColoredPoint struct {
	IPoint
	Color color.RGBA
}

var (
	mu      sync.Mutex //guards mapping
	mapping = make(map[string]string)
)

func Oop() {
	x := IPoint{2.3, 4.5}
	y := IPoint{5.3, 5.5}
	var path Path
	path = append(path, x)
	path = append(path, y)
	fmt.Println(Distance(x, y))
	fmt.Println(x.Distance(y))
	fmt.Println(path.Distance())
	x.ScaleBy(4.5)
	fmt.Println(x)

	colorpoint := ColoredPoint{x, color.RGBA{1, 2, 3, 4}}
	colorpoint.ScaleBy(2)
	fmt.Println(colorpoint)
}

func Distance(p, q IPoint) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p IPoint) Distance(q IPoint) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p *IPoint) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
