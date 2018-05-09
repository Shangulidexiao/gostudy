// 打印命令行参数
package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stdinDup() {
	lines := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		lines[input.Text()]++
	}

	for line, n := range lines {
		if n > 1 {
			fmt.Printf("重复结果：%d\t%s\n", n, line)
		}
	}
}

func fileDup() {
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
