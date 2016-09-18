package fleet

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_Machines(t *testing.T) {
	machinesJson := `{
	  "machines": [
	    {
	      "id": "4b3c49d6f2b14c4e84fb2eea0bdbb0f3",
	      "primaryIP": "172.17.8.102"
	    },
	    {
	      "id": "90eeaaea56c84a98a6c20cc1cd756d0f",
	      "primaryIP": "172.17.8.103"
	    },
	    {
	      "id": "b6c55728e762434e9e56f23612afba1d",
	      "primaryIP": "172.17.8.101"
	    }
	  ]
	}`

	baseUrl, _ := url.ParseRequestURI("http://fleet.example.com:4001/")
	senderMock := &requestSenderMock{
		httpResponse: &http.Response{
			Status:     "OK",
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(machinesJson)),
		},
	}

	client := &Client{
		baseUrl:       baseUrl,
		requestSender: senderMock,
	}

	machines, err := client.Machines()

	if nil != err {
		t.Error("Error supposed to be nil")
	}

	if http.MethodGet != senderMock.httpRequest.Method {
		t.Error("Request header must be GET")
	}

	if "http://fleet.example.com:4001/fleet/v1/machines" != senderMock.httpRequest.URL.String() {
		t.Error("Request URL must be http://fleet.example.com:4001/fleet/v1/machines")
	}

	if "application/json" != senderMock.httpRequest.Header.Get("Content-Type") {
		t.Error("Content-Type header must be application/json")
	}

	if len(machines) != 3 {
		t.Error("Return slice must contain exactly 3 items")
	}

	if machines[0].ID != "4b3c49d6f2b14c4e84fb2eea0bdbb0f3" {
		t.Error("Wrong machine 0 ID")
	}

	if machines[0].PrimaryIP != "172.17.8.102" {
		t.Error("Wrong machine 0 PrimaryIP")
	}

	if machines[2].ID != "b6c55728e762434e9e56f23612afba1d" {
		t.Error("Wrong machine 2 ID")
	}

	if machines[2].PrimaryIP != "172.17.8.101" {
		t.Error("Wrong machine 2 PrimaryIP")
	}
}

func TestClient_MachinesInvalidJsonError(t *testing.T) {
	machinesJson := `{{{Invalid JSON}`

	baseUrl, _ := url.ParseRequestURI("http://fleet.example.com:4001/")
	senderMock := &requestSenderMock{
		httpResponse: &http.Response{
			Status:     "OK",
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(machinesJson)),
		},
	}

	client := &Client{
		baseUrl:       baseUrl,
		requestSender: senderMock,
	}

	machines, err := client.Machines()

	if err == nil {
		t.Error("Error supposed not to be nil")
	}

	if len(machines) > 0 {
		t.Error("Machines supposed to be empty")
	}
}
