package main

import (
	"fmt"
	"time"
	"runtime"
)

func main()  {

	name := "Aaaaa"

	go func( n string ) {
		fmt.Printf( "hello %s", n )
	}( name )

	name = "bbbbb"

	fmt.Println( runtime.NumGoroutine() )

	time.Sleep( time.Millisecond * 100 )
}


