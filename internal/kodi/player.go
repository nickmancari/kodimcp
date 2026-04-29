package kodi

import (
)

func (c *KodiClient) Pause(playerID int) (any, error) {
	res, err := c.Call("Player.PlayPause", map[string]any{
		"playerid": playerID,
	})
	if err != nil {
		return nil, err
	}

	return res["result"], nil
}
