package fleet

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	basePath = "/fleet/v1"
)

// A Client manages communication with the fleet API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// BaseURL for API requests.
	BaseURL *url.URL

	// BasePath for fleet base API endpoint
	BasePath string
}

// NewClient returns a new fleet API client.
func NewClient(baseURLStr string) (*Client, error) {
	baseURL, err := url.Parse(baseURLStr)
	if err != nil {
		return nil, err
	}

	return &Client{
		client:   &http.Client{},
		BaseURL:  baseURL,
		BasePath: basePath,
	}, nil
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	urlStr := c.BaseURL.String() + c.BasePath + endpoint
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// An Error reports an error caused by an API request.
// fleet API docs: https://coreos.com/fleet/docs/latest/api-v1.html#error-communication
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Body     struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Body.Message)
}

// Do sends an API requests and returns the API response.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return resp, nil
}

func CheckResponse(resp *http.Response) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	errResp := &ErrorResponse{Response: resp}
	data, err := ioutil.ReadAll(resp.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errResp)
	}

	return errResp
}
