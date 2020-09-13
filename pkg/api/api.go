package api

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func get(url string) (*resty.Response, error) {
	return call(http.MethodGet, url)
}

func call(method, url string) (*resty.Response, error) {
	// TODO: factor this out and up to the main call
	client := resty.New()
	resp, err := client.R().Execute(method, url)
	if err != nil {
		return nil, fmt.Errorf("could not send request to server: %v", err)
	}

	// check the status code
	if resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("server responded with status code %d: %s", resp.StatusCode(), resp.String())
	}

	return resp, nil
}
