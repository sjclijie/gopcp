package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {

	var ipv4 [4]uint = [...]uint{192, 168, 0, 4}

	fmt.Println(ipv4)

	ipv4_string := []string{"192.168.1.1", "192.168.1.2", "192.168.1.3" }

	fmt.Println(ipv4_string)

	fmt.Println(cap(ipv4_string))
	fmt.Println(ipv4_string[:cap(ipv4_string)])

	ipv4_string = append(ipv4_string, "192.168.3.1")

	fmt.Println(cap(ipv4_string))

	fmt.Println(ipv4_string)

	ips := make([]string, 100)

	fmt.Println(cap(ips))

	ips = append(ips, "192.168.1.1")

	fmt.Println(ips)

	var ipswitches = map[string]bool{}

	ipswitches["192.168.1.1"] = true
	ipswitches["192.168.1.2"] = false


	fmt.Println(ipswitches)

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input your name: ")

	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Printf("Found an error: %s\n", err)
	} else {
		fmt.Printf("Hello, %s\n", input)
	}
}
