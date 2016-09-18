package fleet

import (
	"io"
	"net/http"
	"net/url"
)

const (
	basePath = "/fleet/v1"
)

type Client struct {
	baseUrl       *url.URL
	err           error
	requestSender requestSender
}

// NewClient creates a basic client that is configured to be used
// with the given fleet HTTP API URL.
func NewClient(rawUrl string) *Client {
	baseUrl, err := url.ParseRequestURI(rawUrl)

	if nil != err {
		baseUrl = &url.URL{}
	}

	return &Client{
		baseUrl:       baseUrl,
		err:           err,
		requestSender: &httpRequestSender{http.Client{}},
	}
}

func NewValidClient(rawUrl string) (*Client, error) {
	client := NewClient(rawUrl)

	if nil != client.Error() {
		return nil, client.Error()
	}

	return client, nil
}

func (c *Client) Error() error {
	return c.err
}

func (c *Client) createRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	requestUrl := c.baseUrl
	requestUrl.Path = basePath + endpoint

	request, err := http.NewRequest(method, requestUrl.String(), body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	return request, nil
}
