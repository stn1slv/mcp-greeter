// Package main implements an MCP (Model Context Protocol) server that provides greeting functionality.
// This server demonstrates basic MCP tool implementation patterns.
package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Input represents the input parameters for the greet tool.
type Input struct {
	Name string `json:"name" jsonschema:"the name of the person to greet"`
}

// Output represents the response from the greet tool.
type Output struct {
	Greeting string `json:"greeting" jsonschema:"the greeting to tell to the user"`
}

// SayHi implements the greet tool functionality.
// It takes a name as input and returns a personalized greeting message.
// This function serves as an example of how to implement MCP tool handlers.
func SayHi(ctx context.Context, req *mcp.CallToolRequest, input Input) (
	*mcp.CallToolResult,
	Output,
	error,
) {
	return nil, Output{Greeting: "Ćao " + input.Name}, nil
}

// main initializes and runs the MCP server.
// The server listens on stdin/stdout and provides a greeting tool.
func main() {
	log.Println("Starting MCP server...")

	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v1.0.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "Generate a personalized greeting message"}, SayHi)

	log.Println("MCP server is ready and listening on stdin/stdout")
	log.Println("Available tools: greet")

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
