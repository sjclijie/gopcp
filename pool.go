package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

func main() {

	defer debug.SetGCPercent(1)

	var count int32

	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}

	pool := sync.Pool{New: newFunc}

	//New字段的使用
	v1 := pool.Get()
	fmt.Println("Value 1: ", v1)

	pool.Put(10)
	pool.Put(11)
	pool.Put(13)
	pool.Put(14)

	v2 := pool.Get()
	fmt.Println("Value 2: ", v2)

	pool.New = nil

	v3 := pool.Get()
	fmt.Println("Value 3: ", v3)

	debug.SetGCPercent(100)
	runtime.GC()

	v4 := pool.Get()
	fmt.Println("Value 4:", v4)

}
