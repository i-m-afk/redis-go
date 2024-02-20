package main

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServerConcurrency(t *testing.T) {

	time.Sleep(100 * time.Millisecond) // Wait for the server to start

	// Create multiple connections concurrently
	numClients := 10
	for i := 0; i < numClients; i++ {
		go func(clientNum int) {
			conn, err := net.Dial("tcp", "localhost:6379")
			if err != nil {
				t.Errorf("Failed to connect for client %d: %v", clientNum, err)
				return
			}
			defer conn.Close()

			commands := []string{"ping", "ping"}
			for _, cmd := range commands {
				_, err := conn.Write([]byte(cmd + "\n"))
				if err != nil {
					t.Errorf("Error writing to connection for client %d: %v", clientNum, err)
					return
				}

				response := make([]byte, 1024)
				n, err := conn.Read(response)
				if err != nil {
					t.Errorf("Error reading response for client %d: %v", clientNum, err)
					return
				}

				fmt.Printf("Client %d response: %s\n", clientNum, string(response[:n]))
			}
		}(i)
	}

	// Wait for the connections to finish
	time.Sleep(1 * time.Second)
}
