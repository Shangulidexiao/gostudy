package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

func main() {
	stu := Student {"hanjian",12}
	fmt.Println(stu)
}
