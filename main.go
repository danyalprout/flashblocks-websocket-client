package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
)

func main() {
	var format string

	// Parse command line arguments
	flag.StringVar(&format, "format", "compressed", "Format: 'plain' or 'compressed' (default: compressed)")
	flag.Parse()

	// Get URL as positional argument
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: WebSocket URL is required as first argument\n")
		fmt.Fprintf(os.Stderr, "Usage: %s <websocket_url> [-format <plain|compressed>]\n", os.Args[0])
		os.Exit(1)
	}

	url := args[0]

	// Validate format
	if format != "plain" && format != "compressed" {
		fmt.Fprintf(os.Stderr, "Error: Format must be 'plain' or 'compressed'\n")
		os.Exit(1)
	}

	// Connect to WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket:", err)
	}
	defer conn.Close()

	// Set up signal handling for immediate termination
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Goroutine to read messages
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			processMessage(message, format)
		}
	}()

	// Wait for interrupt signal
	<-interrupt
	os.Exit(0)
}

// processMessage handles the message based on the specified format
func processMessage(message []byte, format string) {
	switch format {
	case "plain":
		fmt.Printf("%s\n", string(message))
	case "compressed":
		// Decompress the message
		decompressed, err := decompressBrotli(message)
		if err != nil {
			return
		}

		// Pretty print as JSON
		var jsonData interface{}
		if err := json.Unmarshal(decompressed, &jsonData); err != nil {
			// If it's not valid JSON, just print the decompressed data
			fmt.Printf("%s\n", string(decompressed))
		} else {
			// Pretty print JSON
			prettyJSON, err := json.MarshalIndent(jsonData, "", "  ")
			if err != nil {
				return
			}
			fmt.Printf("%s\n", string(prettyJSON))
		}
	}
}

// decompressBrotli decompresses brotli-compressed data
func decompressBrotli(data []byte) ([]byte, error) {
	reader := brotli.NewReader(bytes.NewReader(data))

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return decompressed, nil
}
