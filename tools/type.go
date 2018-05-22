package tools

import "fmt"

func PrintDefaultValue() {
	var name string
	var age int
	var score float32
	var success bool
	var mymap map[string]int
	var myslice []string
	var ch chan string
	var p *int = &age

	fmt.Println("字符串初始值:", name)
	fmt.Println("整型初始值:", age)
	fmt.Println("浮点型初始值:", score)
	fmt.Println("布尔型初始值:", success)
	fmt.Println("map初始值:", mymap)
	fmt.Println("slice初始值:", myslice)
	fmt.Println("chan初始值:", ch)
	fmt.Println("指针初始值:", p)
	fmt.Println("指针初始值:", &score)
	*p++
	fmt.Println(*p)
}
