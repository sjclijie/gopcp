package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main() {

	syncChan := make(chan struct{}, 2)

	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}

		fmt.Println("Stoppend. [received]")
		syncChan <- struct{}{}
	}()

	go func() {

		countMap := make(map[string]int)

		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v \n", countMap)
		}

		close(mapChan)

		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan

}
