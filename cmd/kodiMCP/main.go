package main

import (
	"fmt"
	"os"

	"kodimcp/internal/kodi"
	"kodimcp/internal/mcpserver"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	kodiURL := getenv("KODI_URL", "http://0.0.0.0:8080")
	kodiUser := os.Getenv("KODI_USER")
	kodiPass := os.Getenv("KODI_PASSWORD")

	k := kodi.New(kodiURL, kodiUser, kodiPass)

	s := mcpserver.Init(k)

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
