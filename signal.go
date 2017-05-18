package main

import (
	"os"
	"syscall"
	"os/signal"
	"sync"
	"fmt"
	"time"
)

func main() {

	/*
	sigRecv := make(chan os.Signal, 1)

	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT }

	//SIGKILL和 SIGSTOP 对它们的响应只能是执行默认操作

	signal.Notify(sigRecv, sigs...)

	for sig := range sigRecv {
		fmt.Printf("Received a signal: %s\n", sig)
	}

	//关闭
	signal.Stop(sigRecv)
	close(sigRecv)
	*/

	sigRecv1 := make(chan os.Signal, 1)
	sigs1 := []os.Signal{syscall.SIGINT }
	signal.Notify(sigRecv1, sigs1...)

	sigRecv2 := make(chan os.Signal, 1)
	signal.Notify(sigRecv2, syscall.SIGQUIT)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {

		for sig := range sigRecv1 {
			fmt.Printf("Recevied a signal from sigRecv1: %s\n", sig)
		}

		fmt.Printf("End. [sigRecv1]")

		wg.Done()

	}()

	go func() {

		for sig := range sigRecv2 {
			fmt.Printf("Recevied a signal from sigRecv2: %s\n", sig)
		}

		fmt.Printf("End. [sigRecv2]")

		wg.Done()

	}()

	time.Sleep(10 * time.Second)

	signal.Stop( sigRecv1 )

	close( sigRecv1 )

	wg.Wait()
}
