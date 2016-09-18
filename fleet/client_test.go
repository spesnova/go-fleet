package fleet

import (
	"testing"
)

func TestNewClientValid(t *testing.T) {
	client, err := NewValidClient("http://valid-url.com:5000/fleet")

	if nil != err {
		t.Error("Error supposed to be nil")
	}

	if nil == client {
		t.Error("Client supposed not to be nil")
	}
}

func TestNewClientInValid(t *testing.T) {
	client, err := NewValidClient("invalid.url")

	if nil == err {
		t.Error("Error supposed not to be nil")
	}

	if nil != client {
		t.Error("Client supposed to be nil")
	}
}
