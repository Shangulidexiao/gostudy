package ch7

import (
	"fmt"
	"io"
)

func Exec() {
	var b ByteConter
	b.Write([]byte("helloword"))
	fmt.Fprintf(&b, "hello word:%d", 9)
	fmt.Println(b)

	var w WrodConter
	fmt.Fprintf(&w, "Hello world")
	fmt.Println(w)

	var _ io.Writer = &b
}
