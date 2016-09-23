package fleet

import (
	"bytes"
	"encoding/json"
	"errors"
)

const (
	unitsPath = "/units"
)

type Unit struct {
	Name         string        `json:"name,omitempty"`
	Options      []*UnitOption `json:"options,omitempty"`
	DesiredState string        `json:"desiredState,omitempty"`
	CurrentState string        `json:"currentState,omitempty"`
	MachineID    string        `json:"machineID,omitempty"`
}

type UnitOption struct {
	Section string `json:"section,omitempty"`
	Name    string `json:"name,omitempty"`
	Value   string `json:"value,omitempty"`
}

type unitsResponse struct {
	Units []Unit `json:"units,omitempty"`
}

func (c *Client) Units() ([]Unit, error) {
	var unitsResp unitsResponse

	req, err := c.NewRequest("GET", unitsPath, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(req, &unitsResp)
	if err != nil {
		return nil, err
	}

	return unitsResp.Units, nil
}

// Unit returns single unit by name
// fleet API docs: https://github.com/coreos/fleet/blob/master/Documentation/api-v1.md#get-a-unit
func (c *Client) Unit(name string) (*Unit, error) {
	unit := &Unit{}

	req, err := c.NewRequest("GET", unitsPath+"/"+name, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(req, &unit)
	if err != nil {
		return nil, err
	}

	return unit, nil
}

// createOrUpdateUnit creates or updates an unit
func (c *Client) createOrUpdateUnit(u Unit) error {
	switch u.DesiredState {
	case "inactive":
		break
	case "loaded":
		break
	case "launched":
		break
	default:
		return errors.New("Invalid desired state")
	}

	j, err := json.Marshal(u)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("PUT", unitsPath+"/"+u.Name, bytes.NewReader(j))
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Submit(name string, opts []*UnitOption) error {
	unit := Unit{
		Name:         name,
		Options:      opts,
		DesiredState: "inactive",
	}

	return c.createOrUpdateUnit(unit)
}

func (c *Client) Load(name string, opts []*UnitOption) error {
	unit := Unit{
		Name:         name,
		DesiredState: "loaded",
	}

	if len(opts) > 0 {
		unit.Options = opts
	}

	return c.createOrUpdateUnit(unit)
}

// Start sends HTTP request to fleet HTTP API to launch an unit.
// If you want to submit and load and launch at once, pass UnitOption slice to opts.
// Otherwise, set opts to nil.
func (c *Client) Start(name string, opts []*UnitOption) error {
	unit := Unit{
		Name:         name,
		DesiredState: "launched",
	}

	if len(opts) > 0 {
		unit.Options = opts
	}

	return c.createOrUpdateUnit(unit)
}

func (c *Client) Stop(name string) error {
	unit := Unit{
		Name:         name,
		DesiredState: "loaded",
	}

	return c.createOrUpdateUnit(unit)
}

func (c *Client) Unload(name string) error {
	unit := Unit{
		Name:         name,
		DesiredState: "inactive",
	}

	return c.createOrUpdateUnit(unit)
}

func (c *Client) deleteUnit(name string) error {
	req, err := c.NewRequest("DELETE", unitsPath+"/"+name, nil)
	if err != nil {
		return err
	}

	_, err = c.Do(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Destroy(name string) error {
	return c.deleteUnit(name)
}
