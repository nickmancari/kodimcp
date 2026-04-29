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
}
