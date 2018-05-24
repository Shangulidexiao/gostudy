package tools

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func Exec() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println(sum(1, 2, 3, 4, 5))
	values := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(values...))
}

//可变参数
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
