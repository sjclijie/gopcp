package main

import (
	"time"
	"context"
	"fmt"
)

/**
Context的使用

context.TODO()

context.Background()

主要方法
	WithCancel()
	WithDeadline()
	WithTimeout() return WithDeadline( parent, time.Now().Add( timeout ) )
	WithValue
*/

func Add(ctx context.Context) interface{} {

	ctx = context.WithValue(ctx, "hello", "world")
	ctx = context.WithValue(ctx, "aaa", "bbb")

	select {
	case <-ctx.Done():
		return -1
	default:
		return ctx.Value("hello")
	}
}

func main() {

	timeout := time.Second * 3

	ctx, _ := context.WithTimeout(context.Background(), timeout)

	fmt.Println(Add(ctx))

	//自动超时
	{

		timeout := time.Second * 3
		context.WithTimeout(context.Background(), timeout)
		//....
	}

	{
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(2 * time.Second)
			cancel()
		}()
		//....
	}

}
