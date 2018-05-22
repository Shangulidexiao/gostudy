// 打印命令行参数
package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StdinDup() {
	lines := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入多行文字，统计重复的行数")

	for input.Scan() {
		if input.Text() == "" {
			break
		}
		lines[input.Text()]++
	}
	repeat := 0
	for line, n := range lines {
		if n > 1 {
			repeat++
			fmt.Printf("重复结果：%d\t%s\n", n, line)
		}
	}
	if repeat == 0 {
		fmt.Println("没有重复行")
	}
	fmt.Println("\n按回车键退出")
	input.Scan()
}

func FileDup() {
	lines := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, lines)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err)
				continue
			}
			countLines(f, lines)
			f.Close()
		}
	}

	for line, n := range lines {
		if n > 1 {
			fmt.Printf("重复结果：%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, lines map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		lines[strings.TrimSpace(input.Text())]++
	}
}
