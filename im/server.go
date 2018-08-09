package im

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type client chan<- string

var (
	entering = make(chan client)
	messages = make(chan string)
	leaving  = make(chan client)
)

// Server is start a tcp server
func Server() {
	port := flag.Int("port", 8080, "请输入要监听的端口号")
	host := flag.String("host", "localhost", "请输入要监听的端口号")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *host, *port)

	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer server.Close()
	fmt.Println("server is start at:", addr)
	go loopServer()

	for {
		client, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer client.Close()
		go talk(client)
	}
}

func talk(client net.Conn) {
	go read(client)
	go write(client)
}

func write(client net.Conn) {
	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		io.WriteString(client, buf.Text())
	}
}

func read(client net.Conn) {
	buf := bufio.NewScanner(client)
	var clientString string
	for buf.Scan() {
		clientString += buf.Text()
	}
}
func loopServer() {
	clients := make(map[client]bool)
	for {
		select {
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		}
	}
}
