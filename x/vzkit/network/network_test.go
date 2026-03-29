package network

import "testing"

func TestParse(t *testing.T) {
	got, err := Parse("host-only")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if got.Mode != ModeHostOnly {
		t.Fatalf("Mode = %q, want %q", got.Mode, ModeHostOnly)
	}
}

func TestCreateAttachmentNone(t *testing.T) {
	got, err := CreateAttachment(Config{Mode: ModeNone})
	if err != nil {
		t.Fatalf("CreateAttachment: %v", err)
	}
	if got.ID != 0 {
		t.Fatalf("CreateAttachment(ModeNone) returned non-zero attachment")
	}
}

func TestCreateDeviceNone(t *testing.T) {
	got, err := CreateDevice(Config{Mode: ModeNone})
	if err != nil {
		t.Fatalf("CreateDevice: %v", err)
	}
	if got.ID != 0 {
		t.Fatalf("CreateDevice(ModeNone) returned non-zero device")
	}
}
