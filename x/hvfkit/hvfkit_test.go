//go:build darwin && arm64

package hvfkit

import (
	"errors"
	"strings"
	"testing"
	"unsafe"

	"github.com/tmc/apple/hypervisor"
)

func TestCallAndCreateRecoverPanics(t *testing.T) {
	want := errors.New("missing symbol")
	err := call("hv_test", func() hypervisor.HVReturn {
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

func TestReleaseNilHandles(t *testing.T) {
	var cfg Config
	if err := cfg.Release(); err != nil {
		t.Fatalf("Config.Release = %v", err)
	}
	if err := cfg.Close(); err != nil {
		t.Fatalf("Config.Close = %v", err)
	}
	var vcpu VCPUConfig
	if err := vcpu.Release(); err != nil {
		t.Fatalf("VCPUConfig.Release = %v", err)
	}
	var gic GICConfig
	if err := gic.Release(); err != nil {
		t.Fatalf("GICConfig.Release = %v", err)
	}
	if err := osRelease(nil); err != nil {
		t.Fatalf("osRelease(nil) = %v", err)
	}
}
