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

func TestDataMap(t *testing.T) {
	input := []byte("zero-copy map test")
	d := DataCreate(input)
	m := d.Map()
	defer m.Release()
	if m.Len() != len(input) {
		t.Fatalf("DataMap.Len() = %d, want %d", m.Len(), len(input))
	}
	if !bytes.Equal(m.Bytes(), input) {
		t.Fatalf("DataMap.Bytes() = %q, want %q", m.Bytes(), input)
	}
}

func TestDataMapRelease(t *testing.T) {
	d := DataCreate([]byte("release test"))
	m := d.Map()
	if m.Len() == 0 {
		t.Fatal("DataMap should not be empty before release")
	}
	m.Release()
	if m.Len() != 0 {
		t.Fatal("DataMap.Len() should be 0 after release")
	}
	if m.Bytes() != nil {
		t.Fatal("DataMap.Bytes() should be nil after release")
	}
	// Double release should be safe.
	m.Release()
}

func TestDataMapEmpty(t *testing.T) {
	d := DataEmpty()
	m := d.Map()
	if m.Len() != 0 {
		t.Fatalf("DataMap of empty data should have Len() 0, got %d", m.Len())
	}
	m.Release()
}

func TestDataCreateNoCopy(t *testing.T) {
	input := []byte("no-copy send test data")
	d := DataCreateNoCopy(input)
	got := DataToBytes(d)
	if !bytes.Equal(got, input) {
		t.Fatalf("DataCreateNoCopy round-trip: got %q, want %q", got, input)
	}
	if d.Len() != len(input) {
		t.Fatalf("DataCreateNoCopy Len() = %d, want %d", d.Len(), len(input))
	}
}

func TestDataApply(t *testing.T) {
	input := []byte("apply iteration test")
	d := DataCreate(input)
	var totalBytes int
	var collected []byte
	ok := d.Apply(func(region Data, offset int, buf []byte) bool {
		totalBytes += len(buf)
		collected = append(collected, buf...)
		return true
	})
	if !ok {
		t.Fatal("Apply returned false")
	}
	if totalBytes != len(input) {
		t.Fatalf("Apply total bytes = %d, want %d", totalBytes, len(input))
	}
	if !bytes.Equal(collected, input) {
		t.Fatalf("Apply collected = %q, want %q", collected, input)
	}
}

func TestDataApplyEarlyStop(t *testing.T) {
	d := DataCreate([]byte("stop early"))
	calls := 0
	ok := d.Apply(func(_ Data, _ int, _ []byte) bool {
		calls++
		return false // stop after first region
	})
	if ok {
		t.Fatal("Apply should return false when callback returns false")
	}
	if calls != 1 {
		t.Fatalf("Apply should have called fn exactly once, got %d", calls)
	}
}

func TestDataApplyEmpty(t *testing.T) {
	d := DataEmpty()
	calls := 0
	ok := d.Apply(func(_ Data, _ int, _ []byte) bool {
		calls++
		return true
	})
	if !ok {
		t.Fatal("Apply on empty data should return true")
	}
	if calls != 0 {
		t.Fatal("Apply on empty data should not call fn")
	}
}
