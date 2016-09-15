package fleet

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

type UnitStateFilter struct {
	UnitName  string
	MachineID string
}

// UnitStates return all unit states
func (c *Client) UnitStates() ([]UnitState, error) {
	return c.unitStateQuery("")
}

// UnitStateFiltered return unit states according to filter
func (c *Client) UnitStateFiltered(filter *UnitStateFilter) ([]UnitState, error) {
	queryString := ""
	if filter != nil {
		query := url.Values{}
		if "" != filter.UnitName {
			query.Set("unitName", filter.UnitName)
		}

		if "" != filter.MachineID {
			query.Set("machineID", filter.MachineID)
		}

		queryString = "?" + query.Encode()
	}

	return c.unitStateQuery(queryString)
}

func (c *Client) unitStateQuery(queryString string) ([]UnitState, error) {
	var statesRes = UnitStatesResponse{}

	req, err := http.NewRequest("GET", c.URL+basePath+statesPath+queryString, nil)
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

	if res.StatusCode != http.StatusOK {
		return statesRes.UnitStates, errors.New(fmt.Sprintf("Wrong status code: %d %s", res.StatusCode, res.Status))
	}

	err = json.NewDecoder(res.Body).Decode(&statesRes)

	if err != nil {
		return nil, err
	}

	return statesRes.UnitStates, nil
}
