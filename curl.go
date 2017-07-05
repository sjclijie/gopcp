package main

import (
	"os"
	"fmt"
	"net/http"
	"io"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./curl <url>")
		os.Exit(1)
	}
}

func main() {

	r, err := http.Get(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer r.Body.Close()

	io.Copy(os.Stdout, r.Body)
}
