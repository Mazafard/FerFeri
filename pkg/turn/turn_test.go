package turn_test

import (
	"github.com/mazafard/ferferi/pkg/turn"
	"net"
	"testing"
)

func TestNewTurnServer(t *testing.T) {
	config := turn.DefaultConfig()
	server := turn.NewTurnServer(config)
	defer server.Close()

	// Verify that the server was started and is listening on the expected port
	conn, err := net.Dial("udp4", "127.0.0.1:19302")
	if err != nil {
		t.Fatalf("Failed to connect to TURN server: %v", err)
	}
	defer conn.Close()

	// TODO: Add additional tests here
}
