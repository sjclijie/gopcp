package main

import (
	"fmt"
	"sync"
	"math/rand"
	"time"
)

const (
	numberGoroutines = 100
	taskLoad         = 2000
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	fmt.Println(numberGoroutines, taskLoad)

	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	for i := 1; i <= numberGoroutines; i++ {
		go worker(tasks, i)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {

	defer wg.Done()

	for {

		task, ok := <-tasks

		if !ok {
			fmt.Printf("Worker: %d: Shutting donw\n", worker)
			return
		}

		fmt.Printf("Worker: %d: Started %s\n", worker, task)

		sleep := rand.Int63n(100)

		time.Sleep(time.Duration(sleep) * time.Microsecond)

		fmt.Printf("Worker: %d: Completed %s\n", worker, task)
	}
}
