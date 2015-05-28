package fleet

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	unitsPath = "/units"
)

type Unit struct {
	Name         string        `json:"name,omitempty"`
	Options      []*UnitOption `json:"options,omitempty"`
	DesireState  string        `json:"desireState,omitempty"`
	CurrentState string        `json:"currentState,omitempty"`
	MachineID    string        `json:"machineID,omitempty"`
}

type UnitOption struct {
	Section string `json:"section,omitempty"`
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
}

func (c *Client) Units() {}

func (c *Client) Submit(name string, opts []*UnitOption) error {
	unit := Unit{
		Options:     opts,
		DesireState: "inactive",
	}

	j, err := json.Marshal(unit)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.URL+basePath+unitsPath+"/"+name, bytes.NewReader(j))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode == 201 {
		return nil
	}

	return nil
}

func (c *Client) Load(name string)    {}
func (c *Client) Start(name string)   {}
func (c *Client) Stop(name string)    {}
func (c *Client) Unload(name string)  {}
func (c *Client) Destroy(name string) {}
