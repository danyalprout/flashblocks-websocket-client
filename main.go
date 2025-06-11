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
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: WebSocket URL is required as first argument\n")
		fmt.Fprintf(os.Stderr, "Usage: %s <websocket_url>\n", os.Args[0])
		os.Exit(1)
	}

	url := args[0]

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket:", err)
	}
	defer conn.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}
			processMessage(message)
		}
	}()

	<-interrupt
	os.Exit(0)
}

func processMessage(message []byte) {
	var jsonData interface{}
	if err := json.Unmarshal(message, &jsonData); err == nil {
		prettyJSON, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			fmt.Printf("%s\n", string(message))
		} else {
			fmt.Printf("%s\n", string(prettyJSON))
		}
		return
	}

	decompressed, err := decompressBrotli(message)
	if err != nil {
		fmt.Printf("%s\n", string(message))
		return
	}

	if err := json.Unmarshal(decompressed, &jsonData); err == nil {
		prettyJSON, err := json.MarshalIndent(jsonData, "", "  ")
		if err != nil {
			fmt.Printf("%s\n", string(decompressed))
		} else {
			fmt.Printf("%s\n", string(prettyJSON))
		}
	} else {
		fmt.Printf("%s\n", string(decompressed))
	}
}

func decompressBrotli(data []byte) ([]byte, error) {
	reader := brotli.NewReader(bytes.NewReader(data))

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return decompressed, nil
}
