package main

import (
	"os/exec"
	"fmt"
	"bytes"
	"io"
	"os"
)

func main() {

	cmd := exec.Command("echo", "-n", "hello")

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Printf("Error: Can not obtain the stdout pipe %s\n", err)
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error: The command can not be startup: %s\n", err)
		return
	}

	var output bytes.Buffer

	for {

		tempOutput := make([]byte, 5)

		n, err := stdout.Read(tempOutput)

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("Error: Can not read data from pipe: %s\n", err)
			}
		}

		if n > 0 {
			output.Write(tempOutput[:n])
		}
	}

	fmt.Println(output.String(), "-----")

	fmt.Println("=============================================")

	reader, writer, err := os.Pipe()

	if err != nil {
		fmt.Println("Error: Create Pipe failed: %s\n", err)
	}

	fmt.Println( reader, writer )
}
