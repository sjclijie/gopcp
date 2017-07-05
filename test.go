package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {

	intChan := make(chan int, 10)

	intChan <- 111

	elem, ok := <-intChan

	fmt.Println(elem, ok)

	procs := runtime.GOMAXPROCS(8)

	thread := debug.SetMaxThreads(10000)

	fmt.Print(thread, procs)
}
