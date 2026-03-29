package storage

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNSDataRoundTrip(t *testing.T) {
	in := []byte("hello")
	out := NSDataToBytes(NSDataFromBytes(in))
	if string(out) != string(in) {
		t.Fatalf("round trip = %q, want %q", out, in)
	}
}

func TestCreateDiskImage(t *testing.T) {
	path := filepath.Join(t.TempDir(), "disk.img")
	if err := CreateDiskImage(path, 1); err != nil {
		t.Fatalf("CreateDiskImage: %v", err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat: %v", err)
	}
	if info.Size() == 0 {
		t.Fatal("CreateDiskImage created zero-sized file")
	}
}
