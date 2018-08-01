package ch6

import (
	"fmt"
)

func Exec() {
	point1 := Point{1, 2}
	point2 := Point{1, 4}
	path := Path{point1, point2}

	d1 := point1.Distance(point2)
	d2 := path.Distance()

	fmt.Println(d1)
	fmt.Println(d2)

	p1 := &point1
	d3 := p1.Distance(point2)
	fmt.Println(d3)

	//更改path 中第一个点的值
	path.ChangeX(2)
	path1 := &path
	d4 := path1.Distance()
	fmt.Println(d4)

	d5 := Point{3, 4}.Distance(point1)
	fmt.Println(d5)
}
