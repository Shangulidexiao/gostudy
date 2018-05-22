package tools

import (
	"fmt"
	"math"
)

func myfloat() {
	f := float32(1 << 24)
	fmt.Println(f, f+1, f == f+1)

	const e = 2.71828
	fmt.Println(e)

	fmt.Println(.567, 1., .567 == .567)

	fmt.Printf("%b\n", 10)

	fmt.Println(math.Sqrt(-1))
}
