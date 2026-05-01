package kodi

import(
)

// A raw response of all the movies in the kodi database with movieid
func (c *KodiClient) GetMovies() (any, error) {
	res, err := c.Call("VideoLibrary.GetMovies", map[string]any{})
	if err != nil {
		return nil, err
	}

	return res["result"], nil
}

func (c *KodiClient) GetMovieFileFromTitle(title string) (any, error) {
	res, err := c.Call("VideoLibrary.GetMovies", map[string]any{
		"properties": []string{"title", "file", "year"},
		"filter": map[string]any{
			"field":    "title",
			"operator": "is",
			"value":    title,
		},
	})
	if err != nil {
		return nil, err
	}

	return res["result"], nil
}
