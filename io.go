package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"net"
)

type userIP net.IP

func main() {

	filename := os.Args[0]

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(filename)

	fmt.Println(string(data))

	fmt.Println("===================")

	slice := []int{10, 20, 30, 40, 50 }

	new_slice := slice[1:3]

	new_slice = append(new_slice, 1000, 2000)

	fmt.Println(slice, new_slice)

	fmt.Println("====================\n")

	var colors map[string]string

	colors = make(map[string]string)

	colors["Red"] = "#da11111"

	if value, exists := colors["Red"]; exists {
		fmt.Println( value )
	}
}


