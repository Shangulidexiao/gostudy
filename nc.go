package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer conn.Close()
	done := make(chan struct{})
	go func() {
		if _, err = io.Copy(os.Stdout, conn); err != nil {
			fmt.Println("server is die")
		}
		done <- struct{}{}
	}()

	if _, err = io.Copy(conn, os.Stdin); err != nil {
		fmt.Println("server is die")
	}
	<-done
}
