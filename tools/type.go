package tools

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

var a int = 5

func q() (stu Student, num int) {
	stu = Student{"hanjian", 12}
	num = 3
	return
}

func echo(argv ...int) {
	var nums [3]int
	nums[0] = 1
	fmt.Println(nums)

	var s = make([]int, 10)
	s = append(s, 5)
	fmt.Println(s)
}
