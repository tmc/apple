package skylight_test

import (
	"fmt"
	"testing"

	"github.com/tmc/apple/x/skylight"
)

func TestActiveSpace(t *testing.T) {
	space, err := skylight.ActiveSpace()
	if err != nil {
		t.Skipf("ActiveSpace unavailable in current execution environment: %v", err)
	}
	if space == 0 {
		t.Errorf("ActiveSpace returned space 0")
	}
	t.Logf("ActiveSpace: %d", space)
}

func ExampleActiveSpace() {
	space, err := skylight.ActiveSpace()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("active space: %d\n", space)
}

func ExampleFocusWithoutRaise() {
	// Focus target window 1001 owned by process PID 1234 without raising windows
	err := skylight.FocusWithoutRaise(1234, 1001)
	if err != nil {
		fmt.Printf("focus failed: %v\n", err)
		return
	}
	fmt.Println("process focused")
}
