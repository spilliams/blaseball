package remotedata

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spilliams/blaseball/pkg"
)

// BlaseballAPI represents an HTTP request interface for interacting with a
// blaseball-centric JSON API
type BlaseballAPI struct {
	blaseURL  string
	blasePath string
	client    *resty.Client
	logger    *logrus.Logger
}

// NewAPI returns a new API client. If given blaseURL is empty, the official
// one will be used (https://www.blaseball.com/database/).
func NewAPI(blaseURL string, blasePath string, logLevel logrus.Level) pkg.RemoteDataSession {
	if len(blaseURL) == 0 {
		blaseURL = "https://www.blaseball.com/database/"
	}
	l := logrus.StandardLogger()
	l.SetLevel(logLevel)
	return &BlaseballAPI{
		blaseURL:  blaseURL,
		blasePath: blasePath,
		client:    resty.New().SetHeader("User-Agent", "github.com/spilliams/blaseball"),
		logger:    l,
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

	b.logger.Debugf("Calling url %s", fullURL)
	resp, err := b.client.R().Execute(method, fullURL)
	if err != nil {
		e := fmt.Errorf("could not send request to server: %v", err)
		return nil, pkg.NewCodedError(e, http.StatusInternalServerError)
	}

	b.logger.Debugf("Request: %v", resp.Request)
	b.logger.Debugf("Response %d: %s", resp.StatusCode(), resp)

	msg := resp.String()
	// special case for blaseball official API!
	if strings.Index(msg, "<title>Offline for Maintenance</title>") > -1 {
		msg = "Offline for Maintenance"
	}

	// check the status code
	if resp.StatusCode() >= 300 {
		e := fmt.Errorf("server responded with status code %d: %s", resp.StatusCode(), msg)
		return nil, pkg.NewCodedError(e, resp.StatusCode())
	}

	return resp, nil
}

func (b *BlaseballAPI) buildURL(finalPath string, queryParams url.Values) (string, error) {
	build, err := url.Parse(b.blaseURL)
	if err != nil {
		return "", err
	}
	build.Path = path.Join(b.blasePath, finalPath)
	if queryParams != nil {
		build.RawQuery = queryParams.Encode()
	}
	return build.String(), nil
}
