package fleet

import (
	"net/url"
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

type UnitStatesFilter struct {
	MachineID string
	UnitName  string
}

func (c *Client) UnitStates(filter *UnitStatesFilter) ([]UnitState, error) {
	var statesResp UnitStatesResponse

	queryString := ""
	if filter.MachineID != "" || filter.UnitName != "" {
		query := url.Values{}
		if filter.MachineID != "" {
			query.Set("machineID", filter.MachineID)
		}

		if filter.UnitName != "" {
			query.Set("unitName", filter.UnitName)
		}

		queryString = "?" + query.Encode()
	}

	req, err := c.NewRequest("GET", statesPath+queryString, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(req, &statesResp)
	if err != nil {
		return nil, err
	}

	return statesResp.UnitStates, nil
}
