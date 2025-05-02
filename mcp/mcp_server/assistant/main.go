package main

import (
    "fmt"

    "github.com/mark3labs/mcp-go/server"
	"github.com/pranavek/assistant-mcp/tools"
)

func main() {

    s := server.NewMCPServer(
        "Assistant",
        "1.0.0",
        server.WithResourceCapabilities(true, true),
        server.WithLogging(),
        server.WithRecovery(),
    )

	tools.RegisterCaculatorTool(s)

    // Start the server
    if err := server.ServeStdio(s); err != nil {
        fmt.Printf("Server error: %v\n", err)
    }
}