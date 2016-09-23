package fleet

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/facebookgo/ensure"
)

func TestClient_Machines(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/fleet/v1/machines", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
			{
				"machines": [
					{
						"id": "8469aa1db00a423e9f14f47378bb5aef",
						"primaryIP":"172.17.8.101"
					},
					{
						"id": "daab00e07fed481d8ccf145b7affc0c5",
						"primaryIP":"172.17.8.102"
					}
				]
			}
		`)
	})

	machines, err := client.Machines()
	if err != nil {
		t.Errorf("Client.Machines returned error: %v", err)
	}

	actual := machines[0].ID
	expected := "8469aa1db00a423e9f14f47378bb5aef"
	ensure.DeepEqual(t, actual, expected)
}
