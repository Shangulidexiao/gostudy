package ch3

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func ExecString() {

	//原生字符串
	desc := `you
	
	name \b`
	fmt.Println(desc)

	name := "Hello, 世界"
	fmt.Printf("name的值：%s\n", name)
	fmt.Println("name 的字节数：", len(name))
	fmt.Println("name 的utf8码点：", utf8.RuneCountInString(name))
	for i := 0; i < len(name); {
		r, size := utf8.DecodeRuneInString(name[i:])
		fmt.Printf("%d\t%d\t%[2]c\n", i, r)
		i += size
	}

	fmt.Println(strconv.Itoa(23))
}
