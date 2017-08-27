package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
	"sync"
	"net/http"
)

var usage = `Usage: %s [options]
	-n requests 	Number of requests to perform
	-c concurrency 	Number of multiple requests to make at a time
	-s timeout 		Seconds to max. wait for each response
	-m method 		Method Name
`;

var (
	requests    int
	concurrency int
	timeout     int
	method      string
	url         string
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, usage, os.Args[0])
	}

	flag.IntVar(&requests, "n", 1000, "")
	flag.IntVar(&concurrency, "c", 100, "")
	flag.IntVar(&timeout, "s", 10, "")
	flag.StringVar(&method, "m", "GET", "")
	flag.Parse()

	if flag.NArg() != 1 {
		exit("Invalid url.")
	}

	method = strings.ToUpper(method)
	url = flag.Args()[0]

	if requests < 1 || concurrency < 1 {
		exit("-n and -c connot be smaller than 1.")
	}

	if requests < concurrency {
		exit("-n connot be less than -c")
	}

	fmt.Println(requests, concurrency, timeout, method, url)

	w := Work{
		Requests:    requests,
		Concurrency: concurrency,
		Timeout:     timeout,
		Method:      method,
		Url:         url,
	}

	w.Run()
}

func exit(msg string) {
	flag.Usage()
	fmt.Fprintln(os.Stderr, "\n[Error] "+msg)
	os.Exit(1)
}

type Work struct {
	Requests    int
	Concurrency int
	Timeout     int
	Method      string
	Url         string
	start       time.Time
	end         time.Time
	results     chan *Result
}

func (w *Work) Run() {

	w.results = make(chan *Result, w.Requests)

	w.start = time.Now()

	var wg sync.WaitGroup

	wg.Add(w.Concurrency)

	for i := 0; i < w.Concurrency; i++ {
		go func() {
			defer wg.Done()
			w.runWorker(w.Requests / w.Concurrency)
		}()
	}

	wg.Wait()

	w.end = time.Now()

	close(w.results)

	//fmt.Println("Total Time: ", w.end.Sub(w.start))

	w.Print()
}

func (w *Work) runWorker(num int) {

	client := &http.Client{
		Timeout: time.Duration(w.Timeout) * time.Second,
	}

	for i := 0; i < num; i++ {

		httpReq, _ := http.NewRequest(w.Method, w.Url, nil)

		start := time.Now()

		client.Do(httpReq)

		end := time.Now()

		fmt.Println("Request Time: ", end.Sub(start))

		w.results <- &Result{Duration: end.Sub(start)}
	}

}

func (w *Work) Print() {

	sum := 0.0

	num := float64(len(w.results))

	for result := range w.results {
		sum += result.Duration.Seconds()
	}

	rps := int(sum / w.end.Sub(w.start).Seconds())
	tpr := sum / num * 1000

	fmt.Printf("Requests per second: \t%d [#/sec]\n", rps)
	fmt.Printf("Time per request: \t%.3f [ms]", tpr)

}

type Result struct {
	Duration time.Duration
}
