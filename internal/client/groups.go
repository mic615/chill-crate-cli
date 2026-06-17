package client

type Group struct {
	Name string
	ID   string
}

func (c *Client) CreateGroup(name string) (Group, error) {
	var g Group
	err := c.do("POST", "/groups", map[string]string{"name": name}, &g)
	return g, err
}

func (c *Client) ListGroups() ([]Group, error) {
	var groups []Group
	err := c.do("GET", "/groups", nil, &groups)
	return groups, err
}
