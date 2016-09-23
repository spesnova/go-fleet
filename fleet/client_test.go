package fleet

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/facebookgo/ensure"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the fleet client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

// setup sets up a test HTTP server along with a fleet.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested..
func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// fleet client configured to use test server
	baseURL, _ := url.Parse(server.URL)
	client, _ = NewClient(baseURL.String())
}

// teardown closes the test HTTP server.
func tearDown() {
	server.Close()
}

func TestNewClient(t *testing.T) {
	client, err := NewClient("http://example.com:49153")

	ensure.NotNil(t, client)
	ensure.Nil(t, err)
}
