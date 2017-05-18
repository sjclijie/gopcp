package main

import (
	"net"
	"fmt"
	"bufio"
)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:9898")

	if err != nil {
		fmt.Printf("Listen error: %s\n", err)
	}

	conn, err := listener.Accept()

	if err != nil {
		fmt.Printf("Accept error: %s\n", err)
	}

	writer := bufio.NewWriter(conn)

	

}
