package main

import "fmt"

func main() {
	/*声明变量*/
	var a int = 8
	fmt.Println(a)
	/*批量声明变量*/
	var (
		b int = 9
		c int = 10
	)
	fmt.Println(b, c)
	/*简短声明*/
	e := 11
	fmt.Println(e)
	/*垃圾桶变量*/
	_, g := 11, 12
	fmt.Println(g)
	/*bool*/
	boolean := true
	fmt.Println(boolean)
	/*int8 int16 int32 int64 byte uint8 uint16 uint32 uint64*/
	var ab int = 9
	var ac int32 = 19
	/*以下将会报错*/
	//var ad int32 = ab + ac
	fmt.Println(ab, ac)
	/*常量只能是sring，bool int*/
	const MAX int = 1000
	fmt.Println(MAX)
	const (
		A = iota
		B
		C
	)
	fmt.Println(A, B, C)
	/*修改p字符串*/
	hello := "hello"
	helloArr := []rune(hello)
	fmt.Println(hello, helloArr)
	helloArr[0] = 'y'
	fmt.Println(string(helloArr), helloArr)

	/*复数*/
	var cc complex64 = 5 + 5i
	fmt.Println(cc)
	fmt.Println(manyArgv)
}

func manyArgv() {
	var arr []int
	arr[0] = 1
	return arr
}
