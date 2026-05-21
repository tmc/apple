//go:build darwin && arm64

package hvfkit

import (
	"errors"
	"strings"
	"testing"
	"unsafe"
)

func TestCallAndCreateRecoverPanics(t *testing.T) {
	want := errors.New("missing symbol")
	err := call("hv_test", func() int32 {
		panic(want)
	})
	if !errors.Is(err, want) {
		t.Fatalf("call panic error = %v, want wrapping %v", err, want)
	}
	_, err = create("hv_create_test", func() unsafe.Pointer {
		panic("missing")
	})
	if err == nil || !strings.Contains(err.Error(), "hv_create_test: missing") {
		t.Fatalf("create panic error = %v", err)
	}
}
