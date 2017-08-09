package main

import (
	"net/http/cookiejar"
	"golang.org/x/net/publicsuffix"
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

type params struct {
	Name string
	Age  int
}

func main() {

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List })

	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{
		Jar: jar,
	}

	//url := "https://requestb.in/15iu8w41"
	url := "http://127.0.0.1:5555/1.0.2/aaa"

	/*

	url-form

	postValues := url.Values{}
	postValues.Add("id", "ID")
	postValues.Add("pwd", "PWD")
	postValues.Add("s-mode", "0")

	strings.NewReader(postValues.Encode())

	httpReq, err := http.NewRequest(http.MethodPost, url, strings.NewReader(postValues.Encode()))

	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	*/

	params := &params{Name: "lijie", Age: 11}

	params_str, err := json.Marshal(params)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(params_str)

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte( params_str )))

	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Set("User-Agent", "Fill/iPhone-1.0.2")

	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(httpReq)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
