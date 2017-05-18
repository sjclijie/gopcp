package main

import (
	"net"
	"fmt"
	"time"
	"bytes"
	"io"
)

func main() {

	conn, err := net.DialTimeout("tcp", "127.0.0.1:9898", time.Second*10)

	if err != nil {
		fmt.Printf("Dial error %s\n", err)
	}

	var dataBuffer bytes.Buffer

	for {

		b := make([]byte, 10)

		n, err := conn.Read(b)

		if err != nil {
			if err == io.EOF {
				fmt.Printf("The connection is closed.")
				conn.Close()
			} else {
				fmt.Printf("Read Error: %s", err)
			}
			break
		}

		dataBuffer.Write(b[:n])
	}

	fmt.Printf(dataBuffer.String())
}
