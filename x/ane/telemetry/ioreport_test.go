//go:build darwin

package telemetry

import (
	"strings"
	"testing"
)

func TestDRAMField(t *testing.T) {
	var b DRAMBytes
	tests := []struct {
		name string
		want *uint64
	}{
		{"ANE AF RD", &b.Read},
		{"ANE AF WR", &b.Write},
		{"ANE DCS RD", &b.CompressedRead},
		{"ANE DCS WR", &b.CompressedWrite},
		{"ANEXL AF RD", &b.Read},
		{"ANEXL DCS WR", &b.CompressedWrite},
		{"GFX AF RD", nil},
		{"VLane Read", nil},
		{"ANE", nil},
		{"", nil},
	}
	for _, tt := range tests {
		if got := dramField(tt.name, &b); got != tt.want {
			t.Errorf("dramField(%q) = %p, want %p", tt.name, got, tt.want)
		}
	}
}

// TestDRAMChannelsPresent checks that this machine still reports the ANE
// memory-controller byte channels under the names the package matches.
// It skips loudly when IOReport or the channels are unavailable.
func TestDRAMChannelsPresent(t *testing.T) {
	names, err := DRAMChannels()
	if err != nil {
		t.Skipf("SKIP: ANE DRAM channels unreadable on this machine: %v", err)
	}
	if len(names) == 0 {
		t.Skipf("SKIP: %q/%q exists but reports no ANE byte channels — names changed on this macOS version", dramGroup, dramSubgroup)
	}
	t.Logf("ANE DRAM channels: %s", strings.Join(names, ", "))
	var b DRAMBytes
	for _, n := range names {
		if dramField(n, &b) == &b.Read {
			return // a read-bytes rail exists, which is what callers ask for
		}
	}
	t.Errorf("no read-bytes channel among %v", names)
}

// TestDRAMMeterIdle exercises the start/stop path. It does no ANE work,
// so it asserts only that the counters read out; other processes may or
// may not move them, and Stop reports "no bytes" as an error, which is a
// legitimate outcome on an idle engine.
func TestDRAMMeterIdle(t *testing.T) {
	m, err := StartDRAM()
	if err != nil {
		t.Skipf("SKIP: StartDRAM: %v", err)
	}
	b, d, err := m.Stop()
	if err != nil {
		t.Logf("no traffic over %v: %v", d, err)
	}
	t.Logf("rd=%d wr=%d dcs_rd=%d dcs_wr=%d over %v", b.Read, b.Write, b.CompressedRead, b.CompressedWrite, d)
	if _, _, err := m.Stop(); err == nil {
		t.Error("second Stop succeeded, want error")
	}
}

func TestDRAMBytesTotal(t *testing.T) {
	b := DRAMBytes{Read: 1, Write: 2, CompressedRead: 4, CompressedWrite: 8}
	if got := b.Total(); got != 15 {
		t.Errorf("Total() = %d, want 15", got)
	}
	if !b.Available() {
		t.Error("Available() = false, want true")
	}
	if (DRAMBytes{}).Available() {
		t.Error("zero DRAMBytes reports available")
	}
}
