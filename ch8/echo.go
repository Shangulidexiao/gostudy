package ch8

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func echo() {
	port := flag.Int("port", 8080, "请输入正确的端口号")

	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *port)
	server, err := net.Listen("tcp", addr)
	fmt.Println("server is start at :", addr)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer server.Close()
	for {
		conn, err := server.Accept()
		fmt.Printf(" %s is connected\n", conn.RemoteAddr())
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer conn.Close()
		go func() {
			for {
				client := bufio.NewScanner(conn)
				if client.Scan() {
					line := client.Text()
					time.Sleep(1 * time.Second)
					io.WriteString(conn, fmt.Sprintf("%s\n", strings.ToUpper(line)))
					io.WriteString(conn, fmt.Sprintf("%s\n", line))
					io.WriteString(conn, fmt.Sprintf("%s\n", strings.ToLower(line)))
				}
			}
		}()
	}

}
