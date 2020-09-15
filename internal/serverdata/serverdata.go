package serverdata

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spilliams/blaseball/internal"
	"github.com/spilliams/blaseball/pkg"
)

type BlaseballAPI struct {
	blaseURL string
	client   *resty.Client
	*logrus.Logger
}

// TODO: DRY this somehow with pkg.officialdata?
func NewAPI(blaseURL string, logLevel logrus.Level) internal.ServerDataSession {
	if len(blaseURL) == 0 {
		blaseURL = "http://localhost:8080/"
	}
	l := logrus.StandardLogger()
	l.SetLevel(logLevel)
	return &BlaseballAPI{
		blaseURL: blaseURL,
		client:   resty.New().SetHeader("User-Agent", "github.com/spilliams/blaseball"),
		Logger:   l,
	}
}

func (b *BlaseballAPI) get(path string, queryParams url.Values) (*resty.Response, error) {
	return b.call(http.MethodGet, path, queryParams)
}

func (b *BlaseballAPI) call(method, path string, queryParams url.Values) (*resty.Response, error) {
	fullURL, err := b.buildURL(path, queryParams)
	if err != nil {
		return nil, err
	}

	b.Tracef("Calling url %s", fullURL)
	resp, err := b.client.R().Execute(method, fullURL)
	if err != nil {
		e := fmt.Errorf("could not send request to server: %v", err)
		return nil, pkg.NewCodedError(http.StatusInternalServerError, e)
	}

	b.Tracef("Request: %v", resp.Request)
	b.Tracef("Response %d: %s", resp.StatusCode(), resp)

	msg := resp.String()
	// special cases for blaseball official API!
	if strings.Index(msg, "<title>Offline for Maintenance</title>") > -1 {
		msg = "Offline for Maintenance"
	} else if strings.Index(msg, "<!doctype html>") > -1 {
		return nil, pkg.NewCodedErrorf(http.StatusNotAcceptable, "URL %s leads to an html page, not JSON data. Are you querying the base API you mean to be?", fullURL)
	}

	// check the status code
	if resp.StatusCode() >= 300 {
		e := fmt.Errorf("server responded with status code %d: %s", resp.StatusCode(), msg)
		return nil, pkg.NewCodedError(resp.StatusCode(), e)
	}

	return resp, nil
}

func (b *BlaseballAPI) buildURL(path string, queryParams url.Values) (string, error) {
	build, err := url.Parse(b.blaseURL)
	if err != nil {
		return "", err
	}
	build.Path = path
	if queryParams != nil {
		build.RawQuery = queryParams.Encode()
	}
	return build.String(), nil
}

func addShowFKQuery(q map[string][]string, showFK bool) map[string][]string {
	if !showFK {
		return q
	}
	if q == nil {
		q = make(map[string][]string)
	}
	q["showForbiddenKnowledge"] = []string{"true"}
	return q
}
