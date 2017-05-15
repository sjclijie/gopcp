package main

import "fmt"

type Talk interface {
	Hello(userName string) string
	Talk(hard string) (saying string, end bool, err error)
}

func main() {

	var v interface{}

	switch v.(type) {

	case string:
		fmt.Println("The string is %s", v.(string))
	case int, uint, int8, int16, int32, int64:
		fmt.Println("The string is %s", v.(int))
	default:
		fmt.Println("Unknown type")
	}

	ints := []int{10, 9, 8, 7, 6 }

	for i := range ints {
		fmt.Println(i, "====>", ints[i])
	}

	fmt.Println("===============")

	for i := 0; i < 10; i++ {

		defer func(i interface{}) {
			fmt.Println("hello -> ", i)
		}(i)
	}
}
