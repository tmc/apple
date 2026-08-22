package skylight_test

import (
	"testing"
	"unsafe"

	"github.com/tmc/apple/x/skylight"
)

// TestEventRecordOffsets pins every offset named in EventRecord's
// provenance comment. The struct reaches most of these through anonymous
// padding, so a change to any earlier field silently shifts the later ones;
// this test is what makes that loud.
func TestEventRecordOffsets(t *testing.T) {
	var r skylight.EventRecord

	if got := unsafe.Sizeof(r); got != 248 {
		t.Fatalf("sizeof = %d, want 248", got)
	}

	base := uintptr(unsafe.Pointer(&r))
	tests := []struct {
		name string
		addr uintptr
		want uintptr
	}{
		{"DeclaredLength", uintptr(unsafe.Pointer(&r.DeclaredLength)), 0x04},
		{"EventType", uintptr(unsafe.Pointer(&r.EventType)), 0x08},
		{"LocationX", uintptr(unsafe.Pointer(&r.LocationX)), 0x10},
		{"WinLocationX", uintptr(unsafe.Pointer(&r.WinLocationX)), 0x20},
		{"EventTime", uintptr(unsafe.Pointer(&r.EventTime)), 0x30},
		{"EventFlags", uintptr(unsafe.Pointer(&r.EventFlags)), 0x38},
		{"WindowID", uintptr(unsafe.Pointer(&r.WindowID)), 0x3c},
		{"ConnectionID", uintptr(unsafe.Pointer(&r.ConnectionID)), 0x40},
		{"SessionField", uintptr(unsafe.Pointer(&r.SessionField)), 0x48},
		{"Attributes", uintptr(unsafe.Pointer(&r.Attributes)), 0x54},
		{"ActivationState", uintptr(unsafe.Pointer(&r.ActivationState)), 0x8a},
		{"KeyCode", uintptr(unsafe.Pointer(&r.KeyCode)), 0x90},
		{"AppendixPtr", uintptr(unsafe.Pointer(&r.AppendixPtr)), 0xe0},
	}
	for _, tt := range tests {
		if got := tt.addr - base; got != tt.want {
			t.Errorf("%s at %#x, want %#x", tt.name, got, tt.want)
		}
	}
}
