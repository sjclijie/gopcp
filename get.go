package main

import (
	"net/http"
	"os"
	"fmt"
	"io"
)

func main() {

	r, err := http.Get(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Create(os.Args[2])

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	dest := io.MultiWriter(os.Stdout, file)

	io.Copy(dest, r.Body)

	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
