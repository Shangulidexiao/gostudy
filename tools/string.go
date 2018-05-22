package tools

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func StringTest() {
	s := "hello,世界"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	for i, sub := range s {
		fmt.Printf("%d\t%c\n", i, sub)
	}

	fmt.Println(strings.HasPrefix("hello", "e"))
}
