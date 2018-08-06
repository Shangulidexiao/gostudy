package ch7

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

//MySPrintf is test fmt.Fprintf
func MySPrintf(template string, number int) {
	fmt.Fprintf(os.Stdout, template, number)
}

//FReadString  Read File Data
func FReadString(path string) {
	buf := make([]byte, 1024)

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatalf("can't open file :%s", path)
	}
	n, err := file.Read(buf)
	if err != nil {
		log.Fatalf("can't open read file :%d ,%s", n, path)
	}
	str := bytes.NewBuffer(buf).String()
	fmt.Println(str)
}
