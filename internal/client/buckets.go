package client

import "fmt"

type Bucket struct {
	Name string
	ID   string
}

func (c *Client) CreateBucket(name, groupId string) (Bucket, error) {
	var b Bucket
	err := c.do("POST", "/buckets", map[string]string{"name": name, "group_id": groupId}, &b)
	return b, err
}

func (c *Client) ListBuckets(groupID string) ([]Bucket, error) {
	var buckets []Bucket
	err := c.do("GET", fmt.Sprintf("/groups/%s/buckets", groupID), nil, &buckets)
	return buckets, err
}

func (c *Client) DeleteBucket(bucketID string, force bool) error {
	if force {
		return c.do("DELETE", fmt.Sprintf("/buckets/%s?force=true", bucketID), nil, nil)
	}
	return c.do("DELETE", fmt.Sprintf("/buckets/%s", bucketID), nil, nil)
}
