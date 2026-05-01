package mcpserver

import (
	"context"
	"fmt"

	"kodimcp/internal/kodi"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func AddPlayerTools(s *server.MCPServer, k *kodi.KodiClient) {
	s.AddTool(
		mcp.NewTool("kodi_pause",
			mcp.WithDescription("Toggle play/pause for the active Kodi video player"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := k.Pause(1)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Pause toggled: %v", result)), nil
		},
	)
	
	s.AddTool(
		mcp.NewTool("kodi_play_file",
			mcp.WithDescription("Play a file or network path in Kodi"),
			mcp.WithString("file",
				mcp.Required(),
				mcp.Description("File path or URL Kodi can access, such as smb://server/share/movie.mkv"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			file, err := req.RequireString("file")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			result, err := k.PlayFile(file)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Started playback: %v", result)), nil
		},
	)

	s.AddTool(
		mcp.NewTool("kodi_get_now_playing",
			mcp.WithDescription("Get the info on what is now currently playing on the kodi"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := k.GetNowPlaying(1)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Kodi responded: %v", result)), nil
		},
	)
}

func AddStandardTools(s *server.MCPServer, k *kodi.KodiClient) {
	s.AddTool(
		mcp.NewTool("kodi_ping",
			mcp.WithDescription("Check whether Kodi is reachable"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := k.Ping()
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Kodi responded: %v", result)), nil
		},
	)

}

func AddVideoLibraryTools(s *server.MCPServer, k *kodi.KodiClient) {
	s.AddTool(
		mcp.NewTool("kodi_get_movies",
			mcp.WithDescription("Get all the movies on the kodi"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			result, err := k.GetMovies()
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			return mcp.NewToolResultText(fmt.Sprint("Kodi responded: %v", result)), nil
		},
	)

	s.AddTool(
		mcp.NewTool("kodi_get_movie_file_by_title",
			mcp.WithDescription("Get a single movie file by the title provided"),
			mcp.WithString("title",
				mcp.Required(),
				mcp.Description("Title of the movie like Jaws or Star Wars"),
			),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			title, err := req.RequireString("title")
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			result, err := k.GetMovieFileFromTitle(title)
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			return mcp.NewToolResultText(fmt.Sprint("Kodi responded: %v", result)), nil
		},
	)
}
