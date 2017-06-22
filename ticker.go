package main

import (
	"time"
	"fmt"
)

func main() {

	intChan := make(chan int, 1)

	ticker := time.NewTicker(time.Second)

	go func() {
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			}
		}
	}()

	for e := range intChan {
		fmt.Println("received...", e)
	}
}
