//go:build darwin

// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"bytes"
	"runtime"
	"testing"
)

func TestDataRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
	}{
		{"empty/nil", nil},
		{"single byte", []byte{0x42}},
		{"hello", []byte("hello, dispatch")},
		{"binary", []byte{0, 1, 2, 255, 254, 253}},
		{"1MB", bytes.Repeat([]byte("ABCDEFGH"), 128*1024)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DataCreate(tt.input)
			wantLen := len(tt.input)
			if got := d.Len(); got != wantLen {
				t.Errorf("Len() = %d, want %d", got, wantLen)
			}
			got := d.Bytes()
			if !bytes.Equal(got, tt.input) {
				t.Errorf("Bytes() mismatch: got %d bytes, want %d bytes", len(got), len(tt.input))
			}
		})
	}
}

func TestDataEmpty(t *testing.T) {
	d := DataEmpty()
	if d.Len() != 0 {
		t.Errorf("DataEmpty().Len() = %d, want 0", d.Len())
	}
	if got := d.Bytes(); got != nil {
		t.Errorf("DataEmpty().Bytes() = %v, want nil", got)
	}
}

func TestDataConcat(t *testing.T) {
	a := DataCreate([]byte("hello, "))
	b := DataCreate([]byte("world"))
	c := DataCreateConcat(a, b)
	want := []byte("hello, world")
	if c.Len() != len(want) {
		t.Errorf("concat Len() = %d, want %d", c.Len(), len(want))
	}
	if got := c.Bytes(); !bytes.Equal(got, want) {
		t.Errorf("concat Bytes() = %q, want %q", got, want)
	}
}

func TestDataGCPressure(t *testing.T) {
	want := []byte("survive gc")
	d := DataCreate(want)
	runtime.GC()
	runtime.GC()
	if n := DataGetSize(d); n != len(want) {
		t.Fatalf("Len() after GC = %d, want %d", n, len(want))
	}
	got := DataToBytes(d)
	if !bytes.Equal(got, want) {
		t.Errorf("Bytes() after GC = %q, want %q", got, want)
	}
}

func TestDataZeroLength(t *testing.T) {
	d := DataCreate([]byte{})
	if d.Len() != 0 {
		t.Errorf("DataCreate([]byte{}).Len() = %d, want 0", d.Len())
	}
}
