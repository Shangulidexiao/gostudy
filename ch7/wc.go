package ch7

import (
	"bufio"
	"fmt"
	"log"
)

//ByteConter is sum of bytes data
type ByteConter int

//WrodConter is sum of word
type WrodConter int

func (c *ByteConter) Write(b []byte) (int, error) {

	*c += ByteConter(len(b))

	return len(b), nil
}

func (w *WrodConter) Write(b []byte) (int, error) {
	var i = 0
	blen := len(b)
	wordLen := 0
	for i < blen {
		a, t, err := bufio.ScanWords(b, false)
		fmt.Println(a, string(t), err)
		if err != nil {
			log.Fatalf("parse error %d %q", t, err)
		}
		wordLen += a
		i += a
		*w += WrodConter(a)
	}
	return wordLen, nil
}
