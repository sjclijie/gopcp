package main

import (
	"fmt"
	"time"
)

func main() {

	intChan := make(chan int, 1)

	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()

	select {
	case elem := <-intChan:
		fmt.Println(" Hello ", elem)
	case <-time.NewTimer(time.Second * 2).C:
		fmt.Println("timeout")
	}

}
