package client

import (
	"bytes"
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

func (c *Client) do(method, path string, body, out any) error {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, c.baseURL+path, bodyReader)
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	//TODO Authorization: Bearer <token>
	if c.user != "" {
		req.Header.Set("X-Stub-User", c.user)
	}

	response, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 300 {
		var apiErr struct {
			Error string `json:"error"`
		}
		json.NewDecoder(response.Body).Decode(&apiErr)
		if apiErr.Error != "" {
			return fmt.Errorf("%s: %s", response.Status, apiErr.Error)
		}
		return fmt.Errorf("request failed: %s", response.Status)
	}

	if out != nil {
		return json.NewDecoder(response.Body).Decode(out)
	}
	return nil
}
