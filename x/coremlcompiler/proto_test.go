package coremlcompiler

import "testing"

func TestReadBytesRejectsOverlongLength(t *testing.T) {
	// field 1, wire type 2, length 2^64-1.
	data := []byte{0x0a, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01}
	r := newProtoReader(data)
	if _, _, err := r.readTag(); err != nil {
		t.Fatalf("readTag: %v", err)
	}
	if _, err := r.readBytes(); err == nil {
		t.Fatal("readBytes accepted an overlong length, want error")
	}
}
