package main

import (
	"log"
	"os"
	"io/ioutil"
	"io"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {

	file, err := os.OpenFile("./errors.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}

	Trace = log.New(ioutil.Discard, "Trace: ", log.Ldate|log.Lshortfile)
	Info = log.New(os.Stdout, "Info: ", log.Ltime|log.Lshortfile)

	//Warning = log.New(file, "Warn: ", log.Ltime|log.Lshortfile)
	Warning = log.New(io.MultiWriter( os.Stderr, os.Stdout, file ), "Warn: ", log.Ltime|log.Lshortfile)
}

func main() {

	//Info.Println("hello world")
	Warning.Println("1111111111111111111")

	//Error.Println("2222222222222")
}
