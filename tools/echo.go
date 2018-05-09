// 打印命令行参数
package tools

import (
	"fmt"
	"os"
	"strings"
)

func forEcho() {
	var s, step string
	for i := 1; i < len(os.Args); i++ {
		s += step + os.Args[i]
		step = " "
	}
	fmt.Println(s)
}

func rangeEcho() {
	s := ""
	step := ""
	for _, arg := range os.Args[1:] {
		s += step + arg
		step = " "
	}

	fmt.Println(s)
}

func joinEcho() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
