package main

import (
	"context"
	"fmt"
	"os"

	"kodimcp/internal/kodi"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	kodiURL := getenv("KODI_URL", "http://192.168.1.50:8080")
	kodiUser := os.Getenv("KODI_USER")
	kodiPass := os.Getenv("KODI_PASSWORD")

	k := kodi.New(kodiURL, kodiUser, kodiPass)

	s := server.NewMCPServer(
		"Kodi LibreELEC MCP",
		"0.1.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	s.AddTool(
		mcp.NewTool("kodi_ping",
			mcp.WithDescription("Check if Kodi JSON-RPC is reachable"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			res, err := k.Call("JSONRPC.Ping", map[string]any{})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return mcp.NewToolResultText(fmt.Sprintf("Kodi responded: %v", res["result"])), nil
		},
	)

	s.AddTool(
		mcp.NewTool("kodi_get_active_players",
			mcp.WithDescription("Show active Kodi players, such as video or audio"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			res, err := k.Call("Player.GetActivePlayers", map[string]any{})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return mcp.NewToolResultText(fmt.Sprintf("%v", res["result"])), nil
		},
	)

	s.AddTool(
		mcp.NewTool("kodi_pause",
			mcp.WithDescription("Toggle play/pause for the active Kodi video player"),
		),
		func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			res, err := k.Call("Player.PlayPause", map[string]any{
				"playerid": 1,
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}
			return mcp.NewToolResultText(fmt.Sprintf("Pause toggled: %v", res["result"])), nil
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

			res, err := k.Call("Player.Open", map[string]any{
				"item": map[string]any{
					"file": file,
				},
			})
			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			return mcp.NewToolResultText(fmt.Sprintf("Started playback: %v", res["result"])), nil
		},
	)

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "server error: %v\n", err)
		os.Exit(1)
	}
}

func getenv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
