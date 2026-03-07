package vzkit

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestParseDisplaySpec(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    DisplayConfig
		wantErr bool
	}{
		{"4k preset", "4k", DisplayConfig{3840, 2160, 144}, false},
		{"uhd alias", "UHD", DisplayConfig{3840, 2160, 144}, false},
		{"1080p preset", "1080p", DisplayConfig{1920, 1080, 144}, false},
		{"fhd alias", "fhd", DisplayConfig{1920, 1080, 144}, false},
		{"720p preset", "720p", DisplayConfig{1280, 720, 144}, false},
		{"hd alias", "hd", DisplayConfig{1280, 720, 144}, false},
		{"retina preset", "retina", DisplayConfig{2560, 1600, 227}, false},
		{"custom WxH", "2560x1440", DisplayConfig{2560, 1440, 144}, false},
		{"custom WxH@PPI", "1920x1080@72", DisplayConfig{1920, 1080, 72}, false},
		{"minimum valid", "640x480", DisplayConfig{640, 480, 144}, false},
		{"maximum valid", "7680x4320", DisplayConfig{7680, 4320, 144}, false},
		{"width too small", "320x480", DisplayConfig{}, true},
		{"height too small", "640x200", DisplayConfig{}, true},
		{"width too large", "8000x1080", DisplayConfig{}, true},
		{"height too large", "1920x5000", DisplayConfig{}, true},
		{"invalid format", "abc", DisplayConfig{}, true},
		{"invalid PPI", "1920x1080@abc", DisplayConfig{}, true},
		{"no dimensions", "x1080", DisplayConfig{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDisplaySpec(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseDisplaySpec(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("ParseDisplaySpec(%q) = %+v, want %+v", tt.input, got, tt.want)
			}
		})
	}
}

func TestDefaultDisplayConfigs(t *testing.T) {
	mac := DefaultDisplayConfig()
	if mac.Width != 1920 || mac.Height != 1200 || mac.PPI != 144 {
		t.Errorf("DefaultDisplayConfig() = %+v", mac)
	}

	linux := DefaultLinuxDisplayConfig()
	if linux.Width != 1920 || linux.Height != 1080 || linux.PPI != 144 {
		t.Errorf("DefaultLinuxDisplayConfig() = %+v", linux)
	}
}

func TestDisplaySlice(t *testing.T) {
	var s DisplaySlice

	if s.String() != "" {
		t.Errorf("empty DisplaySlice.String() = %q", s.String())
	}

	if err := s.Set("1080p"); err != nil {
		t.Fatalf("Set(1080p): %v", err)
	}
	if err := s.Set("4k"); err != nil {
		t.Fatalf("Set(4k): %v", err)
	}

	if len(s) != 2 {
		t.Fatalf("len = %d, want 2", len(s))
	}
	if s[0] != (DisplayConfig{1920, 1080, 144}) {
		t.Errorf("s[0] = %+v", s[0])
	}
	if s[1] != (DisplayConfig{3840, 2160, 144}) {
		t.Errorf("s[1] = %+v", s[1])
	}

	str := s.String()
	if str != "1920x1080@144,3840x2160@144" {
		t.Errorf("String() = %q", str)
	}

	if err := s.Set("bad"); err == nil {
		t.Error("Set(bad) should fail")
	}
}

func TestCreateDiskImage(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test.img")

	if err := CreateDiskImage(path, 1); err != nil {
		t.Fatalf("CreateDiskImage: %v", err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat: %v", err)
	}

	// Sparse file: logical size should be 1 GB.
	if info.Size() != 1*1024*1024*1024 {
		t.Errorf("size = %d, want %d", info.Size(), 1*1024*1024*1024)
	}

	// Creating again should be a no-op (file exists).
	if err := CreateDiskImage(path, 2); err != nil {
		t.Fatalf("CreateDiskImage(existing): %v", err)
	}
	info2, _ := os.Stat(path)
	if info2.Size() != info.Size() {
		t.Errorf("existing file was modified: %d -> %d", info.Size(), info2.Size())
	}
}

func TestSnapshotManagerListDeleteEmpty(t *testing.T) {
	tmp := t.TempDir()
	mgr := NewSnapshotManager(tmp)

	// List on empty directory returns nil.
	snaps, err := mgr.List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if snaps != nil {
		t.Errorf("expected nil, got %d snapshots", len(snaps))
	}

	// Delete non-existent snapshot returns error.
	if err := mgr.Delete("nope"); err == nil {
		t.Error("Delete(nope) should fail")
	}
}

func TestSnapshotManagerMetadataRoundTrip(t *testing.T) {
	tmp := t.TempDir()
	mgr := NewSnapshotManager(tmp)

	// Simulate saved snapshot files.
	snapDir := filepath.Join(tmp, "snapshots")
	if err := os.MkdirAll(snapDir, 0755); err != nil {
		t.Fatal(err)
	}

	now := time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC)
	meta := SnapshotMeta{
		Name:     "checkpoint1",
		Created:  now,
		Size:     1024,
		VMState:  "paused",
		FilePath: filepath.Join(snapDir, "checkpoint1.vmstate"),
	}
	metaBytes, _ := json.MarshalIndent(meta, "", "  ")
	if err := os.WriteFile(filepath.Join(snapDir, "checkpoint1.json"), metaBytes, 0644); err != nil {
		t.Fatal(err)
	}
	// Create the .vmstate file (empty, just for listing).
	if err := os.WriteFile(filepath.Join(snapDir, "checkpoint1.vmstate"), []byte("state"), 0644); err != nil {
		t.Fatal(err)
	}

	snaps, err := mgr.List()
	if err != nil {
		t.Fatalf("List: %v", err)
	}
	if len(snaps) != 1 {
		t.Fatalf("len = %d, want 1", len(snaps))
	}
	if snaps[0].Name != "checkpoint1" {
		t.Errorf("Name = %q", snaps[0].Name)
	}
	if snaps[0].VMState != "paused" {
		t.Errorf("VMState = %q", snaps[0].VMState)
	}

	// Delete the snapshot.
	if err := mgr.Delete("checkpoint1"); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	snaps, _ = mgr.List()
	if len(snaps) != 0 {
		t.Errorf("after delete: len = %d", len(snaps))
	}
}

func TestSnapshotManagerListOrder(t *testing.T) {
	tmp := t.TempDir()
	mgr := NewSnapshotManager(tmp)
	snapDir := filepath.Join(tmp, "snapshots")
	if err := os.MkdirAll(snapDir, 0755); err != nil {
		t.Fatal(err)
	}

	// Create two snapshots with different creation times.
	for i, name := range []string{"older", "newer"} {
		meta := SnapshotMeta{
			Name:    name,
			Created: time.Date(2025, 1, 1+i, 0, 0, 0, 0, time.UTC),
		}
		metaBytes, _ := json.MarshalIndent(meta, "", "  ")
		os.WriteFile(filepath.Join(snapDir, name+".json"), metaBytes, 0644)
		os.WriteFile(filepath.Join(snapDir, name+".vmstate"), []byte{}, 0644)
	}

	snaps, err := mgr.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(snaps) != 2 {
		t.Fatalf("len = %d", len(snaps))
	}
	// Newest first.
	if snaps[0].Name != "newer" {
		t.Errorf("snaps[0].Name = %q, want newer", snaps[0].Name)
	}
	if snaps[1].Name != "older" {
		t.Errorf("snaps[1].Name = %q, want older", snaps[1].Name)
	}
}

func TestNewQueue(t *testing.T) {
	q := NewQueue("com.test.vzkit")
	if q == nil {
		t.Fatal("NewQueue returned nil")
	}
	if q.Queue().ID() == 0 {
		t.Fatal("underlying dispatch.Queue is zero")
	}

	// Sync should execute and return.
	done := make(chan bool, 1)
	q.Sync(func() {
		done <- true
	})
	select {
	case <-done:
	default:
		t.Fatal("Sync did not execute")
	}
}

func TestQueueAsync(t *testing.T) {
	q := NewQueue("com.test.vzkit.async")
	ch := make(chan int, 1)
	q.Async(func() {
		ch <- 42
	})
	select {
	case v := <-ch:
		if v != 42 {
			t.Errorf("got %d, want 42", v)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("Async work did not complete")
	}
}

func TestWrapQueue(t *testing.T) {
	orig := NewQueue("com.test.wrap.inner")
	wrapped := WrapQueue(orig.Queue())
	if wrapped.Queue() != orig.Queue() {
		t.Error("WrapQueue should preserve the underlying queue")
	}
}

func TestErrorCompletionHandler(t *testing.T) {
	h := NewErrorCompletionHandler()
	if h.Done() {
		t.Error("should not be done before callback")
	}
	if h.Error() != nil {
		t.Error("error should be nil before callback")
	}
	if h.Block() == 0 {
		t.Error("Block() should return non-zero")
	}
}

func TestExtractNSErrorMessageZero(t *testing.T) {
	msg := ExtractNSErrorMessage(0)
	if msg != "" {
		t.Errorf("ExtractNSErrorMessage(0) = %q, want empty", msg)
	}
}

func TestFormatNSErrorDetailedZero(t *testing.T) {
	s := FormatNSErrorDetailed(0)
	if s != "" {
		t.Errorf("FormatNSErrorDetailed(0) = %q, want empty", s)
	}
}

func TestNetworkModeConstants(t *testing.T) {
	if NetworkModeNAT != "nat" {
		t.Errorf("NAT = %q", NetworkModeNAT)
	}
	if NetworkModeBridged != "bridged" {
		t.Errorf("Bridged = %q", NetworkModeBridged)
	}
	if NetworkModeVMNet != "vmnet" {
		t.Errorf("VMNet = %q", NetworkModeVMNet)
	}
	if NetworkModeNone != "none" {
		t.Errorf("None = %q", NetworkModeNone)
	}
}

func TestSnapshotMetaJSON(t *testing.T) {
	now := time.Date(2025, 6, 15, 12, 0, 0, 0, time.UTC)
	meta := SnapshotMeta{
		Name:     "test-snap",
		Created:  now,
		Size:     4096,
		VMState:  "paused",
		FilePath: "/tmp/test.vmstate",
	}

	data, err := json.Marshal(meta)
	if err != nil {
		t.Fatal(err)
	}

	var got SnapshotMeta
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if got.Name != meta.Name {
		t.Errorf("Name = %q", got.Name)
	}
	if !got.Created.Equal(meta.Created) {
		t.Errorf("Created = %v", got.Created)
	}
	if got.Size != meta.Size {
		t.Errorf("Size = %d", got.Size)
	}
	if got.VMState != meta.VMState {
		t.Errorf("VMState = %q", got.VMState)
	}
}
