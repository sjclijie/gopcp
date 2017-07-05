package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:9898"
	DELIMITER      = '\t'
)

func main() {

	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)

	if err != nil {
		printServerLog("Listen Error: %s", err)
		return
	}

	defer listener.Close()

	printServerLog("Got listener for the server: (local address: %s)", listener.Addr())

	for {

		conn, err := listener.Accept()

		if err != nil {
			printServerLog("Accept Error: %s", err)
		}

		printServerLog("Established a connection with a client application. ( remote address: %s )", conn.RemoteAddr())

		go handleConn(conn)
	}

	/*
		conn, err := listener.Accept()

		fmt.Println(conn.LocalAddr().Network())
		fmt.Println(conn.RemoteAddr().Network())

		if err != nil {
			fmt.Printf("Accept error: %s\n", err)
		}

		//writer := bufio.NewWriter(conn)
	*/
}

func handleConn(conn net.Conn) {

	for {

		conn.SetReadDeadline(time.Now().Add(time.Second * 2))

		strReq, err := read(conn)

		if err != nil {

			if err == io.EOF {
				printServerLog("The connection is closed by anthoer side.")
			} else {
				printServerLog("Read Error: %s", err)
			}

			break
		}

		printServerLog("Received request: %s", strReq)

		intReq, err := strToint32(strReq)

		if err != nil {
			n, err := write(conn, err.Error())
			printServerLog("Sent error message ( written %d bytes ): %s", n, err)
			continue
		}

		floatResp := math.Cbrt(float64(intReq))

		respMsg := fmt.Sprintf("the cube root of %d is %f", intReq, floatResp)

		n, err := write(conn, respMsg)

		if err != nil {
			printServerLog("Writ Error: %s", err)
		}

		printServerLog("Sent Response ( written %d bytes ): %s", n, respMsg)
	}
}

func write(conn net.Conn, content string) (n int, err error) {

	var buffer bytes.Buffer

	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)

	return conn.Write(buffer.Bytes())
}

func strToint32(str string) (int32, error) {

	num, err := strconv.ParseInt(str, 10, 0)

	if err != nil {
		return 0, fmt.Errorf("%s is not integer", str)
	}

	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0, fmt.Errorf("%d is not 32-bit integer", num)
	}

	return int32(num), nil
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

func printLog(role string, sn int, format string, args ...interface{}) {

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args...))
}

func printServerLog(format string, args ...interface{}) {
	printLog("Server", 0, format, args...)
}
