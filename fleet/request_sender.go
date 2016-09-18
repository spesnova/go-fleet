package fleet

import "net/http"

type requestSender interface {
	send(req *http.Request) (resp *http.Response, err error)
}

type httpRequestSender struct {
	client http.Client
}

func (rs *httpRequestSender) send(req *http.Request) (resp *http.Response, err error) {
	return rs.client.Do(req)
}

