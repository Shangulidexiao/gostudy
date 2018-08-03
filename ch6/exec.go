package ch6

import (
	"fmt"
	"gostudy/set"
)

func Exec() {
	point1 := Point{1, 2}
	point2 := Point{1, 4}
	path := Path{point1, point2}

	d1 := point1.Distance(point2)
	d2 := path.Distance()

	fmt.Println(d1)
	fmt.Println(d2)

	//用指针调用类型的值
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

	//方法值
	distance := point1.Distance
	fmt.Println(distance(point2))

	//方法表达式
	Distance := Point.Distance
	fmt.Println(Distance(point1, point2))

	fmt.Println("-------------------")

	var s1, s2, s3 set.IntSet

	s1.Add(1)
	s1.Add(2)
	s1.Add(3)
	s2.Add(3)
	s2.Add(4)
	s2.Add(5)
	s3.Add(5)

	fmt.Println("s1：", &s1)
	fmt.Println("s2:", &s2)
	fmt.Println("s3:", &s3)

	s1.UnitWith(&s2)
	fmt.Println("s1 | s2:", &s1)

	s2.IntersectWith(&s3)
	fmt.Println("s2 & s3:", &s2)
}
