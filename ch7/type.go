package ch7

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func typeCC() {
	var w io.Writer

	w = os.Stdout

	f, ok := w.(*os.File)

	b, ok := w.(*bytes.Buffer)

	fmt.Println(f, b, ok)

	iow := w.(io.ReadWriter)
	fmt.Println(iow)
	w = iow
	w = iow.(*os.File)
	_, err := os.Open("dafdafh")
	fmt.Printf("%v\n", err)
	fmt.Println(os.IsNotExist(err))

}
