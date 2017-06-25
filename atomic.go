package main

import (
	"sync/atomic"
	"fmt"
)

func main() {

	var counterVal atomic.Value

	counterVal.Store([]int{1, 2, 3, 4, 5 })

	anthorStore( &counterVal )

	fmt.Println( counterVal.Load() )

}

func anthorStore( counterVal *atomic.Value) {
	counterVal.Store([]int{4, 5, 6, 7 })
}
