// 打印命令行参数
package tools

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func ForEcho() {
	var s, step string
	for i := 1; i < len(os.Args); i++ {
		s += step + os.Args[i]
		step = " "
	}
	fmt.Println(s)
}

func RangeEcho() {
	s := ""
	step := ""
	for _, arg := range os.Args[1:] {
		s += step + arg
		step = " "
	}

	fmt.Println(s)
}

func JoinEcho() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func Echo() {
	//flag.Bool 定义要接受一个bool的命令行参数 第一个名字 第二个默认值 第三个描述
	n := flag.Bool("n", false, "最后要不要输出一个空行")
	step := flag.String("s", " ", "分割的字符串")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *step))
	if !*n {
		fmt.Println()
	}
}
