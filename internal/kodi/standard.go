package kodi

import (
)

func (c *KodiClient) Ping() (any, error) {
	res, err := c.Call("JSONRPC.Ping", map[string]any{})
	if err != nil {
		return nil, err
	}

	return res["result"], nil
}
