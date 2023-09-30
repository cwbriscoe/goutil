package net

import (
	"errors"
	"io"
	"net/http"
	"time"
)

// Client represents an http client
type Client struct {
	client  http.Client
	headers map[string]string
}

// ErrNotFound is returned when the response status is 404 Status Not Found
var ErrNotFound = errors.New("404 not found")

// SetTimeout sets the network timeout of the request
func (c *Client) SetTimeout(d time.Duration) {
	c.client.Timeout = d
}

// SetHeader sets a request header
func (c *Client) SetHeader(k, v string) {
	if c.headers == nil {
		c.headers = make(map[string]string)
	}
	c.headers[k] = v
}

// GetHeader retrieves an already set header
func (c *Client) GetHeader(k string) string {
	if c.headers == nil {
		return ""
	}
	v, ok := c.headers[k]
	if !ok {
		return ""
	}
	return v
}

// ClearHeaders deletes the headers map
func (c *Client) ClearHeaders() {
	c.headers = nil
}

// Fetch will load a network url and return the response in []byte
func (c *Client) Fetch(url string) (*http.Response, []byte, error) {
	defer c.ClearHeaders()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return resp, nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return resp, nil, ErrNotFound
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, nil, err
	}

	return resp, body, err
}
