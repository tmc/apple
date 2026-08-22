package skylight_test

import (
	"os"
	"testing"

	"github.com/tmc/apple/x/skylight"
)

func TestSpaceQueries(t *testing.T) {
	activeSpace, err := skylight.ActiveSpace()
	if err != nil {
		t.Skipf("ActiveSpace failed: %v", err)
	}
	if activeSpace == 0 {
		t.Fatalf("ActiveSpace returned 0")
	}

	t.Logf("ActiveSpace = %d", activeSpace)
}

func TestWindowOwnerAndOffSpace(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in -short mode")
	}

	pid := os.Getpid()
	t.Logf("Current test PID: %d", pid)

	// Test FocusWithoutRaise on self with dummy window ID
	err := skylight.FocusWithoutRaise(pid, 1)
	if err != nil {
		t.Logf("FocusWithoutRaise self: %v", err)
	}

	// Test WithMenuShortcutActivation on self
	err = skylight.WithMenuShortcutActivation(pid, 1, func() error {
		t.Log("Inside WithMenuShortcutActivation action callback")
		return nil
	})
	if err != nil {
		t.Logf("WithMenuShortcutActivation self: %v", err)
	}
}
