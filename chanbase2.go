package main

import (
	"fmt"
	"time"
)

var dataChan = make(chan string, 3)

func main() {

	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go receive(dataChan, syncChan1, syncChan2)
	go sender(dataChan, syncChan1, syncChan2)

	<-syncChan2
	<-syncChan2
}

func receive(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {

	<-syncChan1

	for {

		if elem, ok := <-dataChan; ok {
			fmt.Println("received data ", elem, " [receiver]")
		} else {
			break
		}
	}

	fmt.Println("Stopped. [receiver]")

	syncChan2 <- struct{}{}
}

func sender(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {

	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Sent: ", elem, "[Sender]")

		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Sent sync signal [Sender]")
		}
	}

	fmt.Println("Wait 2 seconds...[Sender]")

	time.Sleep(time.Second * 2)

	close(strChan)

	syncChan2 <- struct{}{}
}
