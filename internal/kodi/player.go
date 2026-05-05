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

func (c *KodiClient) PlayFile(file string) (any, error) {
	res, err := c.Call("Player.Open", map[string]any{
		"item": map[string]any{
			"file": file,
		},
	})
	if err != nil {
		return nil, err
	}

	return res["result"], nil

}

func (c *KodiClient) GetNowPlaying(playerID int) (any, error) {
	res, err := c.Call("Player.GetItem", map[string]any{
		"playerid": playerID,
	})
	if err != nil {
		return nil, err
	}

	return res["result"], nil

}
