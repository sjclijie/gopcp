package main

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	court := make(chan int)

	wg.Add(2)

	go player("aaa", court)
	go player("bbb", court)

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {

	defer wg.Done()

	for {

		ball, ok := <-court

		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("Player %s Missed %d\n", name, n)
			close(court)
		}
		return

		fmt.Printf("Player %s hit %d\n", name, ball)

		ball++

		court <- ball
	}
}
