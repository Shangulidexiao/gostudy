package tools

import (
	"strings"
)

//获取路径中的文件名
func Basename(s string) string {
	var basename string
	for i := 0; i < len(s); i++ {
		if s[i] == '/' || s[i] == '\\' {
			basename = s[i+1:]
		}
	}
	for i := len(basename) - 1; i > 0; i-- {
		if basename[i] == '.' {
			basename = basename[:i]
			break
		}
	}
	return basename
}

func BasenameString(s string) (basename string) {

	namestart := strings.LastIndex(s, "/") + 1

	if nameend := strings.LastIndex(s, "."); nameend == -1 {
		basename = s[namestart:len(s)]
	} else {
		basename = s[namestart:nameend]
	}
	return
}
