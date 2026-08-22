package zerocopy

import (
	"bytes"
	"strings"
	"testing"
)

func TestCheckAlias(t *testing.T) {
	buf := make([]byte, 4096)
	for i := range buf {
		buf[i] = byte(i)
	}
	orig := append([]byte(nil), buf...)
	if err := Check(buf, buf); err != nil {
		t.Fatalf("aliased views: %v", err)
	}
	if err := Check(buf[:100], buf); err != nil {
		t.Fatalf("shorter writer over same memory: %v", err)
	}
	if !bytes.Equal(buf, orig) {
		t.Fatalf("Check did not restore original contents")
	}
}

// TestCheckCopyFails is the package's reason to exist: a deliberately
// copying handoff must fail the probe.
func TestCheckCopyFails(t *testing.T) {
	buf := make([]byte, 4096)
	copied := append([]byte(nil), buf...)
	err := Check(buf, copied)
	if err == nil {
		t.Fatalf("copying reader passed the probe")
	}
	if !strings.Contains(err.Error(), "reader holds a copy") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheckFuncAlias(t *testing.T) {
	buf := make([]byte, 64)
	reads := 0
	err := CheckFunc(buf, func() ([]byte, error) {
		reads++
		return buf, nil
	})
	if err != nil {
		t.Fatalf("aliased read func: %v", err)
	}
	if reads < 2 {
		t.Fatalf("read called %d times, want at least 2", reads)
	}
}

// TestCheckFuncLazySnapshotFails covers the reader that copies once on
// first read and returns the stale snapshot forever after: the sentinel
// probe passes by luck, the restore probe catches it.
func TestCheckFuncLazySnapshotFails(t *testing.T) {
	buf := make([]byte, 64)
	var snap []byte
	err := CheckFunc(buf, func() ([]byte, error) {
		if snap == nil {
			snap = append([]byte(nil), buf...)
		}
		return snap, nil
	})
	if err == nil {
		t.Fatalf("lazy-snapshot reader passed the probe")
	}
}

func TestCheckShortReaderFails(t *testing.T) {
	buf := make([]byte, 64)
	if err := Check(buf, buf[:10]); err == nil {
		t.Fatalf("short reader passed the probe")
	}
}

func TestCheckEmpty(t *testing.T) {
	if err := Check(nil, nil); err == nil {
		t.Fatalf("empty buffer passed the probe")
	}
}

func TestProbeOffsetsSmall(t *testing.T) {
	for n := 1; n <= 8; n++ {
		offs := probeOffsets(n)
		for i, o := range offs {
			if o < 0 || o >= n {
				t.Fatalf("n=%d: offset %d out of range", n, o)
			}
			if i > 0 && o <= offs[i-1] {
				t.Fatalf("n=%d: offsets not strictly increasing: %v", n, offs)
			}
		}
		if offs[len(offs)-1] != n-1 || offs[0] != 0 {
			t.Fatalf("n=%d: offsets must cover both ends: %v", n, offs)
		}
	}
}
