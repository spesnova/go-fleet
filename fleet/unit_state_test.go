package fleet

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/facebookgo/ensure"
)

func TestClient_UnitStates(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/fleet/v1/state", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			{
			  "states": [
			    {
			      "systemdSubState": "running",
			      "systemdLoadState": "loaded",
			      "systemdActiveState": "active",
			      "name": "hello.service",
			      "machineID": "8469aa1db00a423e9f14f47378bb5aef",
			      "hash": "eb866625abd32c86077ec7662af46effac7e40e1"
			    },
			    {
			      "systemdSubState": "running",
			      "systemdLoadState": "loaded",
			      "systemdActiveState": "active",
			      "name": "world.service",
			      "machineID": "8469aa1db00a423e9f14f47378bb5aef",
			      "hash": "eb866625abd32c86077ec7662af46effac7e40e1"
			    }
			  ]
			}
		`)
	})

	filter := UnitStatesFilter{}
	unitStates, err := client.UnitStates(&filter)
	if err != nil {
		t.Errorf("Client.UnitStates returned error: %v", err)
	}

	actual := unitStates[0].SystemdSubState
	expected := "running"
	ensure.DeepEqual(t, actual, expected)
}
