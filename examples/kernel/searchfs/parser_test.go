package main

import (
	"encoding/binary"
	"testing"
	"unsafe"
)

func TestSearchFSLayoutRecords(t *testing.T) {
	if got := unsafe.Sizeof(attrlist{}); got != 24 {
		t.Fatalf("attrlist size = %d, want 24", got)
	}
	if got := unsafe.Sizeof(attrreference{}); got != 8 {
		t.Fatalf("attrreference size = %d, want 8", got)
	}
	if got := unsafe.Sizeof(fssearchblock{}); got != 104 {
		t.Fatalf("fssearchblock size = %d, want 104", got)
	}
	if got := unsafe.Sizeof(fsid{}); got != 8 {
		t.Fatalf("fsid size = %d, want 8", got)
	}
}

func TestNextRecordRejectsMalformedBounds(t *testing.T) {
	valid := make([]byte, 24)
	binary.LittleEndian.PutUint32(valid[:4], 20)
	binary.LittleEndian.PutUint32(valid[4:8], 7)
	binary.LittleEndian.PutUint32(valid[8:12], 9)
	binary.LittleEndian.PutUint64(valid[12:20], 11)
	tests := []struct {
		name string
		buf  []byte
		ok   bool
	}{
		{"valid", valid, true},
		{"short-prefix", valid[:19], false},
		{"short-record", append([]byte(nil), valid[:19]...), false},
		{"length-under-prefix", func() []byte { b := append([]byte(nil), valid...); binary.LittleEndian.PutUint32(b[:4], 19); return b }(), false},
		{"length-over-buffer", func() []byte { b := append([]byte(nil), valid...); binary.LittleEndian.PutUint32(b[:4], 25); return b }(), false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fs, id, length, ok := nextRecord(test.buf)
			if ok != test.ok {
				t.Fatalf("nextRecord ok = %v, want %v", ok, test.ok)
			}
			if test.name == "valid" {
				if fs.Dev != 7 || fs.Val != 9 || id != 11 || length != 20 {
					t.Fatalf("nextRecord = fsid{%d,%d}, id %d, length %d; want fsid{7,9}, id 11, length 20", fs.Dev, fs.Val, id, length)
				}
			}
		})
	}
}

func TestParseResultsStopsAtMalformedRecord(t *testing.T) {
	buf := make([]byte, 20)
	binary.LittleEndian.PutUint32(buf[:4], 100)
	if got := parseResults(buf, 1, api{}); len(got) != 0 {
		t.Fatalf("parseResults returned %v for malformed record", got)
	}
}

func TestParseResultsUsesInclusivePathLength(t *testing.T) {
	buf := make([]byte, 20)
	binary.LittleEndian.PutUint32(buf[:4], 20)
	binary.LittleEndian.PutUint32(buf[4:8], 7)
	binary.LittleEndian.PutUint32(buf[8:12], 9)
	binary.LittleEndian.PutUint64(buf[12:20], 11)
	path := []byte("/tmp/searchfs\x00")
	a := api{getpath: func(dst *byte, size uintptr, _ *fsid, _ uint64) int64 {
		copy((*[4096]byte)(unsafe.Pointer(dst))[:size], path)
		return int64(len(path))
	}}
	got := parseResults(buf, 1, a)
	if len(got) != 1 || got[0] != "/tmp/searchfs" {
		t.Fatalf("parseResults = %q, want /tmp/searchfs without NUL", got)
	}
}

func TestParseResultsHandlesMultipleAndFailedPaths(t *testing.T) {
	first := make([]byte, 20)
	second := make([]byte, 20)
	for _, record := range [][]byte{first, second} {
		binary.LittleEndian.PutUint32(record[:4], 20)
	}
	binary.LittleEndian.PutUint64(first[12:20], 11)
	binary.LittleEndian.PutUint64(second[12:20], 12)
	buf := append(first, second...)
	a := api{getpath: func(dst *byte, size uintptr, _ *fsid, id uint64) int64 {
		if id == 12 {
			return -1
		}
		path := []byte("one\x00")
		copy((*[4096]byte)(unsafe.Pointer(dst))[:size], path)
		return int64(len(path))
	}}
	got := parseResults(buf, 2, a)
	if len(got) != 1 || got[0] != "one" {
		t.Fatalf("parseResults = %q, want one successful path", got)
	}
}
