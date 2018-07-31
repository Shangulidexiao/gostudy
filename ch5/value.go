package ch5

import (
	"fmt"
	"strings"
)

func Exec() {
	test()
}

func pre() {
	fmt.Println("-------")
}
func last() {
	fmt.Println("-------")
}
func paint(pre, last func()) {
	pre()

	fmt.Println("now is now")

	s := func(num int) bool {
		if num%2 == 0 {
			return true
		} else {
			return false
		}
	}

	fmt.Println(s(4))

	last()
}

func str(str string) string {
	return strings.Map(func(r rune) rune { return r + 1 }, str)
}

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
