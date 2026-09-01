package gpureplay

import (
	"context"
	"testing"
	"time"
)

func TestAvailable(t *testing.T) {
	err := Available()
	if err != nil {
		t.Logf("MTLReplayer is not available on this machine: %v", err)
	}
}

func TestReplayValidation(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := Replay(ctx, "/nonexistent/path.gputrace", Options{})
	if err == nil {
		t.Errorf("expected error for nonexistent capture path")
	}
}
