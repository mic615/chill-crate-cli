package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type Client struct {
	baseURL string
	user    string
	http    *http.Client
}

func New() *Client {
	return &Client{
		baseURL: viper.GetString("api_url"),
		user:    viper.GetString("user"),
		http:    &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) doDownload(method, path string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.send(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *Client) doReader(method, path string, body io.Reader, size int64, out any) error {
	req, err := http.NewRequestWithContext(context.Background(), method, c.baseURL+path, body)
	if err != nil {
		return err
	}
	req.ContentLength = size
	req.Header.Set("Content-Type", "application/octet-stream")
	return c.doRequest(req, out)
}

func (c *Client) do(method, path string, body, out any) error {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(context.Background(), method, c.baseURL+path, bodyReader)
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return c.doRequest(req, out)
}

func (c *Client) doRequest(req *http.Request, out any) error {
	resp, err := c.send(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}
	return nil
}

func (c *Client) send(req *http.Request) (*http.Response, error) {
	// TODO Authorization: Bearer <token>
	if c.user != "" {
		req.Header.Set("X-Stub-User", c.user)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		defer resp.Body.Close()
		var apiErr struct {
			Error string `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err == nil && apiErr.Error != "" {
			return nil, fmt.Errorf("%s: %s", resp.Status, apiErr.Error)
		}
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}
	return resp, nil
}
