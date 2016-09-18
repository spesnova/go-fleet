package fleet

import (
	"encoding/json"
)

const (
	machinesPath = "/machines"
)

type Machine struct {
	ID        string            `json:"id,omitempty"`
	PrimaryIP string            `json:"primaryIP,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

type machinesResponse struct {
	Machines []Machine `json:"machines,omitempty"`
}

func (c *Client) Machines() ([]Machine, error) {
	var mRes machinesResponse

	req, err := c.createRequest("GET", machinesPath, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.requestSender.send(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&mRes)
	if err != nil {
		return nil, err
	}

	return mRes.Machines, nil
}
