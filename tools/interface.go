package tools

import (
	"bufio"
	"fmt"
	"strconv"
)

type ByteCounter int
type Word struct {
	Line, Words int
}

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

func (w *Word) Wc(s string) (n int, err error) {
	n, t, err := bufio.ScanWords([]byte(s), false)
	if err != nil {
		fmt.Printf("scan error%q %q", err, t)
		return
	}
	w.Words += n
	return
}

func (w *Word) String() string {
	return strconv.Itoa(w.Words)
}
