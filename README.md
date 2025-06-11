# WebSocket Client Utility

A simple Go utility to connect to WebSocket endpoints and display messages with optional brotli compression support.

## Features

- Connect to any WebSocket endpoint
- Two output formats:
  - **plain**: Display raw message data as-is
  - **compressed**: Decompress brotli data and pretty-print JSON
- Default format is compressed
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
# Connect with compressed format (default)
flashblocks-websocket-client ws://example.com/websocket

# Connect with plain format
flashblocks-websocket-client ws://example.com/websocket -format plain

# Connect with compressed format (explicit)
flashblocks-websocket-client ws://example.com/websocket -format compressed
```

If you built from source:

```bash
# Connect with compressed format (default)
./wsclient ws://example.com/websocket

# Connect with plain format
./wsclient ws://example.com/websocket -format plain
```

### Arguments

- First argument (required): WebSocket URL to connect to
- `-format` (optional): Output format - either `plain` or `compressed` (default: `compressed`)

## Examples

```bash
# Basic usage with compressed JSON output
flashblocks-websocket-client ws://localhost:8080/ws

# Raw data output
flashblocks-websocket-client wss://api.example.com/stream -format plain

# Using local build
./wsclient ws://localhost:8080/ws
./wsclient wss://api.example.com/stream -format plain
```

## Output Formats

### Plain Format
Prints raw WebSocket messages exactly as received.

### Compressed Format (Default)
1. Decompresses brotli-compressed data
2. Attempts to parse as JSON
3. Pretty-prints JSON with proper indentation
4. Falls back to raw decompressed data if not valid JSON

## Dependencies

- `github.com/gorilla/websocket` - WebSocket client library
- `github.com/andybalholm/brotli` - Brotli compression library 