package fleet

import "net/http"

type requestSenderMock struct {
	httpRequest  *http.Request
	httpResponse *http.Response
	err          error
}

func (rs *requestSenderMock) send(req *http.Request) (resp *http.Response, err error) {
	rs.httpRequest = req

	return rs.httpResponse, rs.err
}
