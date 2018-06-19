/*
变量 常量 函数
*/
package ch2

import (
	"fmt"
	"math"
)

//省份证
const id = "140430199301452654"

// 常量
const (
	a = iota
	b
	c
	d
	e
	f
)

var youName string = "liuche"
var myName string
var age int

//执行测试
func ExecType() {
	//组合声明 至少一个是声明加赋值 另一个可以只是赋值
	name, err := "l", 'l'
	name1, err := "l", 'l'
	fmt.Println(name)
	fmt.Println(name1, err)

	fmt.Println(id)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(age)

	fmt.Println(youName)
	fmt.Println(myName)

	var p *string = &youName
	mp := &myName
	fmt.Println(p)
	fmt.Println(mp)
}

func ExecFloat() {
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
}
