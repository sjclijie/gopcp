package main

import (
	"fmt"
)

func main() {

	chanCap := 5
	intChan := make(chan int, chanCap)

	for i := 0; i < chanCap; i++ {
		select {
		case intChan <- 1:
		case intChan <- 2:
		case intChan <- 3:
		}
	}

	for e := range intChan {
		fmt.Println(e)
	}
}
