package fleet

import (
	"encoding/json"
	"net/http"
)

const (
	statesPath = "/state"
)

type UnitState struct {
	Hash               string `json:"hash,omitempty"`
	MachineID          string `json:"machineID,omitempty"`
	Name               string `json:"name,omitempty"`
	SystemdActiveState string `json:"systemdActiveState,omitempty"`
	SystemdLoadState   string `json:"systemdLoadState,omitempty"`
	SystemdSubState    string `json:"systemdSubState,omitempty"`
}

type UnitStatesResponse struct {
	UnitStates []UnitState `json:"states,omitempty"`
}

func (c *Client) UnitStates() ([]UnitState, error) {
	var statesRes UnitStatesResponse

	req, err := http.NewRequest("GET", c.URL+basePath+statesPath, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&statesRes)
	if err != nil {
		return nil, err
	}

	return statesRes.UnitStates, nil
}
