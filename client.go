package pokebuild

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type service struct {
	client *Client
}

type Client struct {
	baseURL   url.URL
	userAgent string

	httpClient *http.Client

	common  service
	Pokemon *pokemonService
}

func NewClient() (*Client, error) {
	httpClient := &http.Client{
		Timeout: 60 * time.Second,
	}

	u, err := url.Parse("https://pokebuildapi.fr/api/v1/")
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseURL:    *u,
		userAgent:  "pokebuildapi-go",
		httpClient: httpClient,
	}
	c.common.client = c
	c.Pokemon = (*pokemonService)(&c.common)

	return c, nil
}

func (c *Client) newRequest(url string) (*http.Request, error) {
	u, err := c.baseURL.Parse(url)
	if err != nil {
		return nil, err
	}

	fmt.Println(u)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)

	return req, nil
}

func (c *Client) do(req *http.Request, v any) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	d := json.NewDecoder(resp.Body)

	err = d.Decode(&v)

	return resp, err
}
