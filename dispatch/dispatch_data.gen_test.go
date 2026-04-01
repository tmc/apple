//go:build darwin

// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"bytes"
	"testing"
)

func TestDataCreate(t *testing.T) {
	input := []byte("hello dispatch data")
	d := DataCreate(input)
	if d.Len() != len(input) {
		t.Fatalf("Len() = %d, want %d", d.Len(), len(input))
	}
	got := d.Bytes()
	if !bytes.Equal(got, input) {
		t.Fatalf("Bytes() = %q, want %q", got, input)
	}
}

func TestDataEmpty(t *testing.T) {
	d := DataEmpty()
	if d.Len() != 0 {
		t.Fatalf("DataEmpty().Len() = %d, want 0", d.Len())
	}
}

func TestDataCreateConcat(t *testing.T) {
	a := DataCreate([]byte("hello "))
	b := DataCreate([]byte("world"))
	c := DataCreateConcat(a, b)
	want := []byte("hello world")
	if c.Len() != len(want) {
		t.Fatalf("concat Len() = %d, want %d", c.Len(), len(want))
	}
	if !bytes.Equal(c.Bytes(), want) {
		t.Fatalf("concat Bytes() = %q, want %q", c.Bytes(), want)
	}
}

func TestDataCreateEmpty(t *testing.T) {
	d := DataCreate([]byte{})
	if d.Len() != 0 {
		t.Fatalf("DataCreate(nil).Len() = %d, want 0", d.Len())
	}
}

func TestDataToBytes(t *testing.T) {
	input := []byte("test DataToBytes")
	d := DataCreate(input)
	got := DataToBytes(d)
	if !bytes.Equal(got, input) {
		t.Fatalf("DataToBytes() = %q, want %q", got, input)
	}
	// Verify it's a copy (mutating original doesn't affect extracted).
	got[0] = 'X'
	got2 := DataToBytes(d)
	if got2[0] != 't' {
		t.Fatal("DataToBytes returned a reference, not a copy")
	}
}
