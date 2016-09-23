package fleet

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
	var mResp machinesResponse

	req, err := c.NewRequest("GET", machinesPath, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(req, &mResp)
	if err != nil {
		return nil, err
	}

	return mResp.Machines, nil
}
