package mcpserver

import (

	"kodimcp/internal/kodi"

	"github.com/mark3labs/mcp-go/server"
)

func Init(k *kodi.KodiClient) *server.MCPServer {


	s := server.NewMCPServer(
		"Kodi MCP",
		"0.1.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)


	AddPlayerTools(s, k)
	AddStandardTools(s, k)
	AddVideoLibraryTools(s, k)

	return s
}
