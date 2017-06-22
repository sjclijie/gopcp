package main

import (
	"time"
	"fmt"
)

func main() {

	timer := time.NewTimer(time.Second * 5)

	fmt.Println("Current: ", time.Now());
	expirtaionTime := <-timer.C

	fmt.Println("Expire time: ", expirtaionTime)

	fmt.Println( timer.Stop() )

}
