# WebSocket Client Utility

A simple Go utility to connect to WebSocket endpoints and automatically detect and display plain JSON or brotli-compressed JSON messages.

## Features

- Connect to any WebSocket endpoint
- Automatic format detection:
  - First tries to parse as plain JSON
  - If that fails, attempts brotli decompression then JSON parsing
  - Falls back to raw output if all parsing fails
- Pretty-printed JSON output
- Graceful shutdown with Ctrl+C

## Installation

### Install from GitHub

```bash
go install github.com/danyalprout/flashblocks-websocket-client@latest
```

### Build from source

```bash
git clone https://github.com/danyalprout/flashblocks-websocket-client.git
cd flashblocks-websocket-client
go build -o wsclient
```

## Usage

```bash
# Connect to websocket (auto-detects format)
flashblocks-websocket-client ws://example.com/websocket
```

If you built from source:

```bash
# Connect to websocket (auto-detects format)
./wsclient ws://example.com/websocket
```

### Arguments

- First argument (required): WebSocket URL to connect to

## Examples

```bash
# Basic usage with auto-detection
flashblocks-websocket-client ws://localhost:8080/ws

# Secure websocket
flashblocks-websocket-client wss://api.example.com/stream

# Using local build
./wsclient ws://localhost:8080/ws
./wsclient wss://api.example.com/stream
```

## How It Works

The utility automatically detects the data format for each message:

1. **Plain JSON**: If the raw message can be parsed as JSON, it's pretty-printed directly
2. **Brotli-compressed JSON**: If JSON parsing fails, the message is decompressed with brotli and then parsed as JSON
3. **Raw output**: If both JSON parsing and brotli decompression fail, the raw message is displayed

This means you don't need to specify the format - the tool handles both plain and compressed data automatically.

## Dependencies

- `github.com/gorilla/websocket` - WebSocket client library
- `github.com/andybalholm/brotli` - Brotli compression library 