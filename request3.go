package main

import (
	"context"
	"net/http"
	"io"
	"time"
	"fmt"
)

func DoRequest(ctx context.Context, req *http.Request) (*http.Response, error) {

	client := &http.Client{}

	resp, err := client.Do(req.WithContext(ctx))

	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
	}

	return resp, err
}

func Get(ctx context.Context, url string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	return DoRequest(ctx, req)
}

func Post(ctx context.Context, url string, bodyType string, body io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", bodyType)

	return DoRequest(ctx, req)
}

func PostForm(ctx context.Context, url string, bodyType string, body io.Reader) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodPost, url, body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", bodyType)

	return DoRequest(ctx, req)
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	url := "http://api.suiyueyule.com/1.0.2/feed/new"

	resp, err := Get(ctx, url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Body)

}
