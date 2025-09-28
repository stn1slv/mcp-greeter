// Package main implements an MCP (Model Context Protocol) server that provides greeting functionality.
// This server demonstrates basic MCP tool implementation patterns and prompt resources.
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

// GreetingPrompt implements a prompt resource that generates greeting prompts.
// It provides a template for creating personalized greeting messages.
func GreetingPrompt(ctx context.Context, req *mcp.GetPromptRequest) (
	*mcp.GetPromptResult,
	error,
) {
	// Extract arguments from the request if provided
	var name string = "there"

	if req.Params.Arguments != nil {
		if nameArg, ok := req.Params.Arguments["name"]; ok {
			name = nameArg
		}
	}

	content := "You are a friendly assistant with access to a 'greet' tool that can generate personalized greeting messages. " +
		"When users want to greet someone, suggest using the greet tool with the person's name. " +
		"For example, you can use the greet tool by calling it with: {\"name\": \"" + name + "\"}. " +
		"Generate a warm response for " + name + " that explains how they can use the greet tool. " +
		"Make your explanation welcoming and suitable for both casual and business contexts."

	return &mcp.GetPromptResult{
		Description: "A prompt that teaches how to use the greet tool for " + name,
		Messages: []*mcp.PromptMessage{
			{
				Role: "user",
				Content: &mcp.TextContent{
					Text: content,
				},
			},
		},
	}, nil
}

// main initializes and runs the MCP server.
// The server listens on stdin/stdout and provides a greeting tool and prompt resource.
func main() {
	log.Println("Starting MCP server...")

	// Create a server with a tool and prompt resource.
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v0.0.1"}, nil)

	// Add the greeting tool
	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "Generate a personalized greeting message"}, SayHi)

	// Add the greeting prompt resource
	server.AddPrompt(&mcp.Prompt{
		Name:        "greeting_prompt",
		Description: "A prompt that explains how to use the greet tool effectively",
	}, GreetingPrompt)

	log.Println("MCP server is ready and listening on stdin/stdout")
	log.Println("Available tools: greet")
	log.Println("Available prompts: greeting_prompt")

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
