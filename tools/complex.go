package tools

import "fmt"

func ComplexTest() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)
	fmt.Println(x, y, x+y, x*y)
}
