package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Print("Please input your name: ")

	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
		os.Exit(1)
	}

	name := input[:len(input)-1]

	fmt.Printf("Hello %s, What can i do for you?\n", name)

	for {

		input, err := inputReader.ReadString('\n')

		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			continue
		}

		input = input[:len(input)-1]

		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println( "Bye!" )
			os.Exit(0)
		default:
			fmt.Println( "Sorry, i didn't catch you." )
		}
	}

}
