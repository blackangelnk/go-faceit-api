package faceit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	defaultTimeout = 5 * time.Second
	faceitHost     = "https://open.faceit.com/data/v4"
)

type Client struct {
	Client *http.Client
	apiKey string
	host   string
}

type StatusError struct {
	Err        error
	StatusCode int
}

type pagination struct {
	Offset int
	Limit  int
}

type timestamps struct {
	From int
	To   int
}

type rPagination struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type rTime struct {
	From int `json:"from"`
	To   int `json:"to"`
}

func (p *pagination) toValues(v *url.Values) {
	if p.Limit != 0 {
		v.Add("limit", strconv.Itoa(p.Limit))
	}
	if p.Offset != 0 {
		v.Add("offset", strconv.Itoa(p.Offset))
	}
}

func (t *timestamps) toValues(v *url.Values) {
	if t.From != 0 {
		v.Add("from", strconv.Itoa(t.From))
	}
	if t.To != 0 {
		v.Add("to", strconv.Itoa(t.From))
	}
}

func (r *StatusError) Error() string {
	return r.Err.Error()
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		host:   faceitHost,
		Client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (c *Client) sendRequest(url string, result interface{}) error {
	req, err := http.NewRequest("GET", c.host+url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return &StatusError{
			StatusCode: resp.StatusCode,
			Err:        fmt.Errorf("status code error %d", resp.StatusCode),
		}
	}

	return json.NewDecoder(resp.Body).Decode(&result)
}
