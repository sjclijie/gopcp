package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
	"net/http/cookiejar"
	"time"
	"context"
	"errors"
)

type result struct {
	Status int      `json:"status"`
	Msg    string `json:"msg"`
	Error  string `json:"error"`
	Data   interface{} `json:"data"`
}

type params struct {
	Name string
	Age  int
}

const API = "http://127.0.0.1:5555/1.0.2/aaa"

func main() {

	var result result

	params := &params{Name: "lijie", Age: 22 }

	params_str, err := json.Marshal(params)

	if err != nil {
		fmt.Println(err)
		return
	}

	//req, err := http.NewRequest(http.MethodPost, API, strings.NewReader(string(params_str)))
	req, err := http.NewRequest(http.MethodPost, API, bytes.NewReader(params_str))

	if err != nil {
		fmt.Println(err)
		return
	}

	//timeout
	timeout := 5 * time.Second
	//timeout := 5 * time.Millisecond

	err = doRequest(timeout, req, func(response *http.Response, err error) error {

		if err != nil {
			return err
		}

		defer response.Body.Close()

		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			return err
		}

		fmt.Println(result.Msg, result.Status, result.Data)

		return nil
	})

	fmt.Println(err)
}

func doRequest(timeout time.Duration, req *http.Request, f func(response *http.Response, err error) error) error {

	jar, err := cookiejar.New(nil)

	if err != nil {
		return nil
	}

	/*
	timer := make(chan struct{})

	time.AfterFunc(timeout, func() {
		timer <- struct{}{}
	})
	*/

	//ctx, cancel := context.WithCancel(context.TODO())
	//
	//defer cancel()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req = req.WithContext(ctx)

	/*
	tr := &http.Transport{Dial: func(network, addr string) (net.Conn, error) {
		return net.DialTimeout(network, addr, timeout)
	}}
	*/

	tr := &http.Transport{}

	client := &http.Client{Jar: jar, Transport: tr }

	c := make(chan error, 1)

	go func() {
		c <- f(client.Do(req))
	}()

	select {
	case <-ctx.Done():
		return errors.New("timeout....")
	case err := <-c:
		return err
	}

	/*
	select {
	case <-timer:
		tr.CancelRequest(req)
		return errors.New("timeout....")
	case err := <-c:
		return err
	}
	*/
}
