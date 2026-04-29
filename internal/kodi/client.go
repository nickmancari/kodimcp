package kodi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type KodiClient struct {
	BaseURL  string
	User     string
	Password string
	client   *http.Client
}

func New(baseURL, user, password string) *KodiClient {
	return &KodiClient{
		BaseURL: baseURL,
		User: user,
		Password: password,
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *KodiClient) Call(method string, params any) (map[string]any, error) {
	body := map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
		"params":  params,
	}

	b, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", c.BaseURL+"/jsonrpc", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if c.User != "" {
		req.SetBasicAuth(c.User, c.Password)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	if out["error"] != nil {
		return nil, fmt.Errorf("kodi error: %v", out["error"])
	}

	return out, nil
}
