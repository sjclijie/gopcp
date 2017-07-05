package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:9898"
	DELIMITER      = '\t'
	CLIENT_NUMS    = 1
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go request(1)

	time.Sleep(time.Second * 2)

	go request(2)

	wg.Wait()
}

func printLog(role string, sn int, format string, args ...interface{}) {

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args...))
}

func printClientLog(sn int, format string, args ...interface{}) {
	printLog("Client", sn, format, args...)
}

func write(conn net.Conn, content string) (int, error) {

	var buffer bytes.Buffer

	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)

	return conn.Write(buffer.Bytes())
}

func read(conn net.Conn) (string, error) {

	readBytes := make([]byte, 1)

	var buffer bytes.Buffer

	for {

		_, err := conn.Read(readBytes)

		if err != nil {
			return "", err
		}

		readByte := readBytes[0]

		if readByte == DELIMITER {
			break
		}

		buffer.WriteByte(readByte)
	}

	return buffer.String(), nil
}

func request(id int) {

	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, time.Second*10)

	if err != nil {
		printClientLog(id, "Dial error %s\n", err)
	}

	defer conn.Close()

	printClientLog(id, "Connected to server.( remote address: %s, local address: %s)", conn.RemoteAddr(), conn.LocalAddr())

	for i := 0; i < CLIENT_NUMS; i++ {

		reqId := rand.Int31()

		n, err := write(conn, fmt.Sprintf("%d", reqId))

		if err != nil {
			printClientLog(id, "Writer Error: ", err)
			continue
		}

		printClientLog(id, "Sent request ( written %d bytes ): %d.", n, reqId)

		time.Sleep(200 * time.Microsecond)
	}

	for i := 0; i < CLIENT_NUMS; i++ {

		strResp, err := read(conn)

		if err != nil {
			if err == io.EOF {
				printClientLog(id, "The connection is closed by another side")
			} else {
				printClientLog(id, "Read Error: %s", err)
			}
		}

		printClientLog(id, "Response: %s", strResp)
	}

	wg.Done()
}
