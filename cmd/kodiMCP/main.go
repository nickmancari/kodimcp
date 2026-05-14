package main

import (
	"fmt"
	"kodimcp/internal/kodi"
	"kodimcp/internal/mcpserver"
	"os"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	kodiURL := getenv("KODI_URL", "http://192.168.1.50:8080")
	kodiUser := os.Getenv("KODI_USER")
	kodiPass := os.Getenv("KODI_PASSWORD")

	k := kodi.New(kodiURL, kodiUser, kodiPass)

	s := mcpserver.Init(k)

	httpServer := server.NewStreamableHTTPServer(s)

	addr := getenv("MCP_ADDR", ":8081")

	fmt.Fprintf(os.Stderr, "Kodi MCP listening on %s\n", addr)

	if err := httpServer.Start(addr); err != nil {
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
