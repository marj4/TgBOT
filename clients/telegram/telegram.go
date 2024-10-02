package main

import (
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	host     string
	basePath string // host.com/bot <token>
	http     http.Client
}

// Create client
func New(host, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		http:     http.Client{},
	}

}

// Create basePath
func newBasePath(token string) string {
	return "bot" + token
}

// Implement methood Update()
func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}

	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	//do request( finish it )

}

func (c *Client) doRequest(methood string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   c.basePath + "methood", //finish it
	}

}

func (c *Client) SendUpdates() {

}
