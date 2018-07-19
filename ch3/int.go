/**
第三章练习
*/
package ch3

import "fmt"

//整型测试
func ExecInt() {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)

	var i int8 = 127
	fmt.Println(i, i+1, i*i)

	fmt.Println(4 > 3)

	fmt.Println(1 & 1)
	fmt.Println(0 | 1)
	fmt.Println(3 ^ 1)

	fmt.Println(1 >> 1)
	fmt.Println(1 << 3)
	fmt.Println(5 &^ 1)
	//按位清空 0000 1111 &^ 1111 0000 ybit为1 值为0 ybit为0值为x的值

}
