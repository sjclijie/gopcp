package main

import "fmt"

func main() {

	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}

	close(intChan)

	syncChan := make(chan struct{}, 1)

	go func() {

	Loop:
		for {
			select {
			case e, ok := <-intChan:
				if !ok {
					fmt.Print("End.")
					goto Loop
				}

				fmt.Println("Received: ", e)
			}
		}

		syncChan <- struct{}{}
	}()

	<-syncChan
}
