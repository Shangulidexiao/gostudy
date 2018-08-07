package ch7

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func ivalue() {

	var w io.Writer //接口值由两部分组成：接口类型 接口值 w：type：nil value:nil
	w = os.Stdout   // w: type:*os.File value: 指向标准输出的指针
	w.Write([]byte("hello world\n"))
	w = new(bytes.Buffer) // w: type:*bytes.Buffer value: 指向
	w = nil

	var x interface{} = []int{1, 2, 3, 4}
	x = []string{"a", "b"}
	fmt.Printf("%T\n", x) //真的可以写中文了

	ff()
}
