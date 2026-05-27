package serial

import "testing"

import vz "github.com/tmc/apple/virtualization"

func TestAvailableUnknown(t *testing.T) {
	if Available("nope") {
		t.Fatal("unknown serial kind is available")
	}
}

func TestNewRequiresAttachment(t *testing.T) {
	if _, err := New(PL011, vz.VZSerialPortAttachment{}); err == nil {
		t.Fatal("New with nil attachment error = nil")
	}
}
