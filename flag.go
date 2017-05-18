package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "default name", "name example...")
}

func main() {

	flag.Parse()

	fmt.Println(name)

	fmt.Printf("pid: %d, ppid: %d", os.Getpid(), os.Getppid())
}
