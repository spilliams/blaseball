package api

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/spilliams/blaseball/internal"
)

type BlaseballAPI struct {
	blase  string
	client *resty.Client
}

func NewAPI(base string) internal.RemoteDataSession {
	if len(base) == 0 {
		base = "https://www.blaseball.com/database/"
	}
	return &BlaseballAPI{
		blase:  base,
		client: resty.New(),
	}
}

func (b *BlaseballAPI) get(url string) (*resty.Response, error) {
	return b.call(http.MethodGet, url)
}

func (b *BlaseballAPI) call(method, url string) (*resty.Response, error) {
	url = fmt.Sprintf("%s%s", b.blase, url)
	resp, err := b.client.R().Execute(method, url)
	if err != nil {
		return nil, fmt.Errorf("could not send request to server: %v", err)
	}

	// check the status code
	if resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("server responded with status code %d: %s", resp.StatusCode(), resp.String())
	}

	return resp, nil
}
