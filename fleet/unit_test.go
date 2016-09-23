package fleet

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/facebookgo/ensure"
)

func TestClient_Units(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/fleet/v1/units", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			{
			  "units": [
			    {
			      "options": [
			        {
			          "value": "Useless infinite loop",
			          "section": "Unit",
			          "name": "Description"
			        },
			        {
			          "value": "/bin/bash -c \"while true; do echo 'hello' && sleep 1; done\"",
			          "section": "Service",
			          "name": "ExecStart"
			        },
			        {
			          "value": "multi-user.target",
			          "section": "Install",
			          "name": "WantedBy"
			        }
			      ],
			      "name": "hello.service",
			      "machineID": "8469aa1db00a423e9f14f47378bb5aef",
			      "desiredState": "launched",
			      "currentState": "launched"
			    },
			    {
			      "options": [
			        {
			          "value": "Useless infinite loop",
			          "section": "Unit",
			          "name": "Description"
			        },
			        {
			          "value": "/bin/bash -c \"while true; do echo 'hello' && sleep 1; done\"",
			          "section": "Service",
			          "name": "ExecStart"
			        },
			        {
			          "value": "multi-user.target",
			          "section": "Install",
			          "name": "WantedBy"
			        }
			      ],
			      "name": "world.service",
			      "machineID": "8469aa1db00a423e9f14f47378bb5aef",
			      "desiredState": "launched",
			      "currentState": "launched"
			    }
			  ]
			}
		`)
	})

	units, err := client.Units()
	if err != nil {
		t.Errorf("Client.Units returned error: %v", err)
	}

	actual := units[0].Name
	expected := "hello.service"
	ensure.DeepEqual(t, actual, expected)
}
