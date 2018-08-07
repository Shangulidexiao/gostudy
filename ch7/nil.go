package ch7

import (
	"bytes"
	"io"
)

const debug = false

func ff() {
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf)
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}

/*
接口值：由两部分组成，接口类型和接口值
包含nil指针的接口不是nil接口
buf: type: ByteConter value: nil
nil: type: nil value: nil
声明时要用接口声明
*/
