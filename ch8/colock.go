package ch8

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func colock() {
	port := flag.Int("port", 8080, "端口号必须为整数")
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *port)
	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("colock server is start at：", addr)
	defer server.Close()
	for {
		conn, err := server.Accept()
		fmt.Println("a client is connect")
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer conn.Close()
		go func() {
			for {
				_, err = io.WriteString(conn, time.Now().Format("15:05:05\n"))
				if err != nil {
					fmt.Println("a client is deconnect")
					return
				}
				time.Sleep(1 * time.Second)
			}
		}()
	}
}
