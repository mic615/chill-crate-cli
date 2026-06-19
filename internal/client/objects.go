package client

import (
	"fmt"
	"io"
)

type Object struct {
	FileName string
	ID       string
	Version  int
	Size     int64
}

func (c *Client) DownloadObject(bucketID, fileName string) (io.ReadCloser, error) {
	return c.doDownload("GET", fmt.Sprintf("/buckets/%s/objects/%s", bucketID, fileName))
}

func (c *Client) UploadObject(
	bucketID, fileName string,
	body io.Reader,
	size int64,
) (Object, error) {
	var object Object
	err := c.doReader(
		"POST",
		fmt.Sprintf("/buckets/%s/objects/%s", bucketID, fileName),
		body,
		size,
		&object,
	)
	return object, err
}

func (c *Client) ListObjects(bucketID string) ([]Object, error) {
	var objects []Object
	err := c.do("GET", fmt.Sprintf("/buckets/%s/objects", bucketID), nil, &objects)
	return objects, err
}
