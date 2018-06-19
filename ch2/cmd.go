package ch2

import (
	"flag"
	"fmt"
	"os"
)

func ExecCmd() {
	n := flag.String("name", "", "请输入姓名")
	flag.Parse()
	fmt.Println('-', *n)
	fmt.Println(os.Args)

	fmt.Printf("%s %[1]s %#x", "hanjian", 10)
}
