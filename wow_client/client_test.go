package wow_client

import (
	"github.com/denisskin/word-of-wisdom/wow_server"
	"testing"
	"time"
)

func init() {
	go wow_server.Start(8080, 1000, 10)
	time.Sleep(1 * time.Millisecond)
}

func TestClient_Get(t *testing.T) {
	client := New(":8080")

	message, err := client.Get()

	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if len(message) == 0 {
		t.Errorf("Empty result")
	}
}
