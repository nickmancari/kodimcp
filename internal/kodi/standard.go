package kodi

func (c *KodiClient) Ping() (any, error) {
	res, err := c.Call("JSONRPC.Ping", map[string]any{})
	if err != nil {
		return nil, err
	}

	return res["result"], nil
}

func (c *KodiClient) Shutdown() (any, error) {
	res, err := c.Call("System.Shutdown", map[string]any{})
	if err != nil {
		return nil, err
	}

	return res["result"], nil
}

func (c *KodiClient) Reboot() (any, error) {
	res, err := c.Call("System.Reboot", map[string]any{})
	if err != nil {
		return nil, err
	}

	return res["result"], nil
}
