# mcp-greeter

A tiny example MCP (Model Context Protocol) server written in Go that exposes a single tool, `greet`, which generates short personalized greeting messages.

This repository demonstrates a minimal MCP tool implementation using the `github.com/modelcontextprotocol/go-sdk/mcp` SDK.

## What it does

- Provides a `greet` tool which accepts a JSON input `{ "name": "..." }` and returns a greeting string.
- Provides a prompt resource `greeting_prompt` that describes how to use the `greet` tool.

The handler in `main.go` currently returns greetings using the form `Ćao <name>`.

## Files of interest

- `main.go` – server implementation and handlers (tool + prompt resource).
- `go.mod`, `go.sum` – Go module files.

## Requirements

- Go 1.20+ (or compatible)

## Build

From the repository root:

```bash
go build -o mcp-greeter main.go
```

This produces a single executable named `mcp-greeter`.

## Run

The MCP server implemented here communicates over stdin/stdout (a common pattern for MCP-style tool processes). Run the server:

```bash
./mcp-greeter
```

You should see logs indicating the server is ready and which tools/prompts are available.

## Example: greet tool input / output

The `greet` tool accepts an `Input` JSON object with a single field `name` and returns an `Output` object with a `greeting` field.

Example request (conceptual):

{
  "tool": "greet",
  "input": { "name": "Stas" }
}

Example response (from the tool handler):

{
  "greeting": "Ćao Stas"
}

Note: The exact transport envelope depends on the MCP client you use. This repository implements the server side; use an MCP-capable client (or a test harness) to send properly framed MCP messages over stdin/stdout.

## Customize

- To change the greeting text or locale, edit the `SayHi` function in `main.go`.
- To add more tools, follow the pattern used for `greet` and register them with `mcp.AddTool`.

## Tests

This repo doesn't include unit tests. For small changes, add table-driven tests for tool handlers in a `_test.go` file and run `go test ./...`.

## Contributing

Contributions are welcome. Open an issue or a pull request with a short description of the change.

## License

The project includes a `LICENSE` file in the repository root. Follow the terms specified there.
