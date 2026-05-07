package kodi

import (
	"fmt"
)

// A raw response of all the movies in the kodi database with movieid
func (c *KodiClient) GetMovies() (any, error) {
	res, err := c.Call("VideoLibrary.GetMovies", map[string]any{})
	if err != nil {
		return nil, err
	}

	result, ok := res["result"].(map[string]any)
	if !ok {
		return nil, fmt.Errorf("unexpected result shape")
	}

	return result["movies"].([]any), nil
}

func (c *KodiClient) GetMovieFileFromTitle(title string) (any, error) {
	res, err := c.Call("VideoLibrary.GetMovies", map[string]any{
		"properties": []string{"title", "file", "year"},
		"filter": map[string]any{
			"field":    "title",
			"operator": "contains",
			"value":    title,
		},
	})
	if err != nil {
		return nil, err
	}

	result := res["result"].(map[string]any)
	movies := result["movies"].([]any)

	if len(movies) == 0 {
		return "", fmt.Errorf("no movie found with title %q", title)
	}

	if len(movies) > 1 {
		return "", fmt.Errorf("multiple movies found with title %q", title)
	}

	movie := movies[0].(map[string]any)
	file, ok := movie["file"].(string)
	if !ok || file == "" {
		return "", fmt.Errorf("movie found, but no file path returned")
	}

	return file, nil

}
